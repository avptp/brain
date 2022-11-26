package request

import (
	"context"

	"github.com/avptp/brain/internal/generated/data"
)

type ViewerCtxKey struct{}

func ViewerFromCtx(ctx context.Context) *data.Person {
	v := ctx.Value(ViewerCtxKey{})

	if v == nil {
		return nil
	}

	return v.(*data.Person)
}
