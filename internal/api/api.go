package api

import (
	bs "inventory_cqrs/internal/bootstrap"
	"inventory_cqrs/internal/config"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	config config.Config
	logger *zap.Logger
	container *bs.Container
}

type RouterBinder struct {

}

func NewAPI(cfg config.Config, logger *zap.Logger, container *bs.Container) *http.Server {
	api := &App{
		config: cfg,
		logger: logger,
		container: container,
	}

    srv := &http.Server{
        Addr:    cfg.HTTP.Address,
        Handler: api.bindRoutes(),
    }

	return srv
}

func (s *App) bindRoutes() http.Handler {
	router := chi.NewRouter()
	
	router.Route("/api/v1", func(r chi.Router) {
		s.BindProductsRoutes(r)
	})

	return router
}