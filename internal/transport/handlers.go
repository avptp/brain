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

	mux.Handle("/ping", healthHandler(ctn.GetLogger()))
	mux.Handle("/", graphHandler(ctn))

	return mux
}

func healthHandler(log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")

		_, err := w.Write([]byte("pong"))

		if err != nil {
			log.Error(
				err.Error(),
			)

			http.Error(w, reporting.ErrInternal.Message, http.StatusInternalServerError)
		}
	})
}

func graphHandler(ctn *container.Container) http.Handler {
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
		middleware.NewSetViewer(data),
	)
}
