package request

import (
	"context"

	"github.com/avptp/brain/internal/generated/container"
)

type ContainerCtxKey struct{}

func ContainerFromCtx(ctx context.Context) (ctn *container.Container) {
	ctn, ok := ctx.Value(ContainerCtxKey{}).(*container.Container)

	if !ok {
		panic("trying to access the context container without having previously set it") // unrecoverable situation
	}

	return
}
