package reporting

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime/debug"

	"entgo.io/ent/privacy"
	"github.com/99designs/gqlgen/graphql"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/stoewer/go-strcase"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var constraintRegexp = regexp.MustCompile(`violates .* constraint "([\w]+)"`)
var validationRegexp = regexp.MustCompile(`field "([\w]+).([\w.]+)"`)

func PanicHandler(ctx context.Context, err any) error {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr)
	debug.PrintStack()

	return ErrInternal
}

func NewErrorPresenter(cfg *config.Config) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		switch {
		case data.IsConstraintError(err):
			fields := constraintRegexp.FindStringSubmatch(err.Error())

			err = ErrConstraint

			if len(fields) >= 2 {
				err.Extensions["field"] = fields[1]
			}
		case
			errors.Is(err, types.ErrIDSize),
			errors.Is(err, types.ErrIDInvalidCharacters),
			errors.Is(err, types.ErrIDScanType):
			err = ErrInput
		case data.IsNotFound(err):
			err = ErrNotFound
		case errors.Is(err, privacy.Deny):
			err = ErrUnauthorized
		case data.IsValidationError(err):
			fields := validationRegexp.FindStringSubmatch(err.Error())

			err = ErrValidation

			if len(fields) >= 3 {
				err.Extensions["field"] = fmt.Sprintf(
					"%s.%s",
					fields[1],
					strcase.LowerCamelCase(fields[2]),
				)
			}
		}

		if _, ok := err.Extensions["code"]; !ok {
			if !cfg.Debug {
				return ErrInternal
			}
		}

		return err
	}
}
