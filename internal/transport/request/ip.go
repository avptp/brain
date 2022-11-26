package request

import (
	"context"
)

type IPCtxKey struct{}

func IPFromCtx(ctx context.Context) (ip string) {
	ip, ok := ctx.Value(IPCtxKey{}).(string)

	if !ok {
		panic("trying to access the context IP without having previously set it") // unrecoverable situation
	}

	return
}
