package reporting

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"entgo.io/ent/privacy"
	"github.com/99designs/gqlgen/graphql"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/oklog/ulid/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

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
			err = ErrConstraint
		case
			errors.Is(err, ulid.ErrDataSize),
			errors.Is(err, ulid.ErrInvalidCharacters),
			errors.Is(err, ulid.ErrBufferSize),
			errors.Is(err, ulid.ErrBigTime),
			errors.Is(err, ulid.ErrOverflow),
			errors.Is(err, ulid.ErrMonotonicOverflow),
			errors.Is(err, ulid.ErrScanValue):
			err = ErrInput
		case data.IsNotFound(err):
			err = ErrNotFound
		case errors.Is(err, privacy.Deny):
			err = ErrUnauthorized
		case data.IsValidationError(err):
			err = ErrValidation
		}

		if _, ok := err.Extensions["code"]; !ok {
			if !cfg.Debug {
				return ErrInternal
			}
		}

		return err
	}
}
