package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"water-delivery/internal/config"
	"water-delivery/internal/platform/db"
	"water-delivery/internal/platform/logger"
	postgresRepository "water-delivery/internal/repository/postgres"
	"water-delivery/internal/service"
	httptransport "water-delivery/internal/transport/http"
	"water-delivery/internal/transport/http/handlers"

	"go.uber.org/zap"
)

func New(ctx context.Context) (*http.Server, func() error, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("load config: %w", err)
	}

	log, err := logger.New(cfg.App.LogLevel)
	if err != nil {
		return nil, nil, fmt.Errorf("create logger: %w", err)
	}

	pgCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := db.NewPostgresPool(pgCtx, cfg.Postgres)
	if err != nil {
		_ = log.Sync()
		return nil, nil, fmt.Errorf("connect postgres: %w", err)
	}

	productRepository := postgresRepository.NewProductRepository(pool)
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	router := httptransport.NewRouter(httptransport.Dependencies{
		Config:         cfg,
		Logger:         log,
		ProductHandler: productHandler,
	})

	server := &http.Server{
		Addr:              ":" + cfg.App.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	cleanup := func() error {
		pool.Close()
		return log.Sync()
	}

	log.Info("application initialized",
		zap.String("app", cfg.App.Name),
		zap.String("env", cfg.App.Env),
		zap.String("port", cfg.App.Port),
	)

	return server, cleanup, nil
}
