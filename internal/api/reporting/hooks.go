package reporting

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"entgo.io/ent/privacy"
	"github.com/99designs/gqlgen/graphql"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func PanicHandler(ctx context.Context, err any) error {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr)
	debug.PrintStack()

	return ErrInternal
}

func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)

	switch {
	case data.IsConstraintError(err):
		err = ErrConstraint
	case data.IsNotFound(err):
		err = ErrNotFound
	case errors.Is(err, privacy.Deny):
		err = ErrUnauthorized
	case data.IsValidationError(err):
		err = ErrValidation
	}

	if _, ok := err.Extensions["code"]; !ok {
		return ErrInternal
	}

	return err
}
