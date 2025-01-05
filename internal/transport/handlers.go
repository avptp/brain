package transport

import (
	"log/slog"
	"net/http"

	"entgo.io/contrib/entgql"
	gqlgen "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/api"
	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/transport/middleware"
)

func Mux(ctn *container.Container) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", GraphHandler(ctn))
	mux.Handle("/stripe", ctn.GetBiller().WebhookHandler())
	mux.Handle("/ping", HealthHandler(ctn.GetLogger()))

	return mux
}

func GraphHandler(ctn *container.Container) http.Handler {
	cfg := ctn.GetConfig()
	data := ctn.GetData()

	// Initialize GraphQL handler
	handler := gqlgen.NewDefaultServer(
		api.NewExecutableSchema(api.Config{
			Resolvers: ctn.GetResolver(),
		}),
	)

	// Configure handler
	handler.SetRecoverFunc(reporting.PanicHandler)
	handler.SetErrorPresenter(reporting.NewErrorPresenter(cfg))
	handler.Use(extension.FixedComplexityLimit(100))
	handler.Use(entgql.Transactioner{TxOpener: data})

	// Chain middlewares
	return middleware.Chain(handler,
		middleware.NewSetIP(ctn.GetIpStrategy()),
		middleware.NewSetAuth(data),
	)
}

func HealthHandler(log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "text/plain")

		_, err := res.Write([]byte("pong"))

		if err != nil {
			log.Error(
				err.Error(),
			)

			http.Error(res, reporting.ErrInternal.Message, http.StatusInternalServerError)
		}
	})
}
