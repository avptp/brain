package jobs

import (
	"context"
	"time"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/madflojo/tasks"
)

func cleanExpiredAuthorizations(ctx context.Context, ctn *container.Container) *tasks.Task {
	return &tasks.Task{
		Interval: time.Hour,
		TaskFunc: func() error {
			cfg := ctn.GetConfig()
			log := ctn.GetLogger()
			data := ctn.GetData()
			allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

			vertices, err := data.Authorization.
				Delete().
				Where(
					authorization.CreatedAtLT(
						time.Now().Add(-cfg.AuthorizationMaxAge),
					),
				).
				Exec(allowCtx)

			if err != nil {
				return err
			}

			log.Info(
				"task completed: clean expired authorizations",
				"vertices", vertices,
			)

			return nil
		},
		ErrFunc: func(e error) {
			log := ctn.GetLogger()

			log.Error(
				e.Error(),
			)
		},
	}
}
