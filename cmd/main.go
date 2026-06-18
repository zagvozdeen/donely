package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zagvozdeen/donely/internal/api"
	"github.com/zagvozdeen/donely/internal/config"
	"github.com/zagvozdeen/donely/internal/db"
	"github.com/zagvozdeen/donely/internal/logger"
	"github.com/zagvozdeen/donely/internal/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to run application", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	log := logger.New()

	pool, err := db.New(ctx, cfg, log)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	storage := store.New(pool)

	return api.NewApplication(cfg, log, storage).Run(ctx)
}
