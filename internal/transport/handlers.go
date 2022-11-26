package transport

import (
	"net/http"

	"entgo.io/contrib/entgql"
	gqlgen "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/api/resolvers"
	"github.com/avptp/brain/internal/generated/api"
	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/transport/middleware"
	"go.uber.org/zap"
)

func Mux(ctn *container.Container) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/ping", healthHandler(ctn.GetLogger()))
	mux.Handle("/", graphHandler(ctn))

	return mux
}

func healthHandler(log *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")

		_, err := w.Write([]byte("pong"))

		if err != nil {
			log.Error(err)
			http.Error(w, reporting.ErrInternal.Message, http.StatusInternalServerError)
		}
	})
}

func graphHandler(ctn *container.Container) http.Handler {
	d := ctn.GetData()

	// Initialize GraphQL handler
	handler := gqlgen.NewDefaultServer(
		api.NewExecutableSchema(api.Config{
			Resolvers: resolvers.NewResolver(
				d,
			),
		}),
	)

	// Configure handler
	handler.SetRecoverFunc(reporting.PanicHandler)
	handler.SetErrorPresenter(reporting.ErrorPresenter)
	handler.Use(extension.FixedComplexityLimit(100))
	handler.Use(entgql.Transactioner{TxOpener: d})

	// Chain middlewares
	return middleware.Chain(handler,
		middleware.NewSetContainer(ctn),
		middleware.NewSetIP(ctn.GetIpStrategy()),
		middleware.NewSetViewer(d),
	)
}
