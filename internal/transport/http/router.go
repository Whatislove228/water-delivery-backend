package httptransport

import (
	"net/http"

	"water-delivery/internal/config"
	"water-delivery/internal/transport/http/handlers"
	"water-delivery/internal/transport/http/middleware"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Dependencies struct {
	Config         config.Config
	Logger         *zap.Logger
	ProductHandler *handlers.ProductHandler
}

func NewRouter(deps Dependencies) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recover(deps.Logger))
	r.Use(middleware.Logging(deps.Logger))

	healthHandler := handlers.NewHealthHandler(deps.Config)
	r.Get("/health", healthHandler.ServeHTTP)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/products", deps.ProductHandler.List)
	})

	return r
}
