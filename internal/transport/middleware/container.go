package middleware

import (
	"context"
	"net/http"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/transport/request"
)

func NewSetContainer(ctn *container.Container) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			subCtn, err := ctn.SubContainer()

			if err != nil {
				panic(err) // unrecoverable situation
			}

			defer func() {
				err := subCtn.Delete()

				if err != nil {
					ctn.GetLogger().Error(
						err.Error(),
					)
				}
			}()

			ctx := context.WithValue(
				r.Context(),
				request.ContainerCtxKey{},
				subCtn,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
