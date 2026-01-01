package main

import (
	"context"
	bs "inventory_cqrs/internal/bootstrap"
	"inventory_cqrs/internal/config"
	loggerify "inventory_cqrs/internal/observability"
	"inventory_cqrs/internal/outbox"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	logger, err := loggerify.New(loggerify.Config{Service: "api", Env: cfg.Env, Level: cfg.Logger.Level})

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	container := bs.NewWorkerContainer(cfg, logger)

	outboxWorker := outbox.NewOutboxWorker(
		container, logger, cfg.Worker)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go outboxWorker.Run()

	<-ctx.Done()

	outboxWorker.Shutdown()

	logger.Info("worker stopped")
}
