package jobs

import (
	"context"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/madflojo/tasks"
)

type Job func(ctx context.Context, ctn *container.Container) *tasks.Task

var All = []Job{
	cleanExpiredAuthorizations,
}
