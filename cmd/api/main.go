package main

import (
	"context"
	"inventory_cqrs/internal/api"
	bs "inventory_cqrs/internal/bootstrap"
	"inventory_cqrs/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	loggerify "inventory_cqrs/internal/observability"

	"go.uber.org/zap"
)

func main() {
    cfg := config.Load()

    logger, err := loggerify.New(loggerify.Config{Service: "api", Env: cfg.Env, Level: cfg.Logger.Level})

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	container := bs.InjectServices(cfg)

    srv := api.NewAPI(cfg, logger, container)

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go func() {
        logger.Info("HTTP server started", zap.String("addr", cfg.HTTP.Address))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("server error", zap.Error(err))
        }
    }()


    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    logger.Info("shutting down HTTP server")
    srv.Shutdown(shutdownCtx)
}