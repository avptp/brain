package request

import (
	"context"

	"github.com/avptp/brain/internal/generated/data"
)

type AuthnCtxKey struct{}

func AuthnFromCtx(ctx context.Context) *data.Authentication {
	v := ctx.Value(AuthnCtxKey{})

	if v == nil {
		return nil
	}

	return v.(*data.Authentication)
}
