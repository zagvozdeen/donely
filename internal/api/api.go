package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/zagvozdeen/donely/internal/config"
	"github.com/zagvozdeen/donely/internal/logger"
	"github.com/zagvozdeen/donely/internal/store"
)

type Application struct {
	cfg       *config.Config
	log       *logger.Logger
	store     *store.Store
	validator *validator.Validate
}

func NewApplication(cfg *config.Config, log *logger.Logger, store *store.Store) *Application {
	return &Application{
		cfg:       cfg,
		log:       log,
		store:     store,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (a *Application) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:        ":8000",
		Handler:     a.getMux(),
		ErrorLog:    a.log.GetLog(),
		ReadTimeout: 30 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	go func() {
		select {
		case err := <-errCh:
			a.log.Error("Failed to listen and serve server", err)
		case <-time.After(time.Millisecond * 500):
			a.log.Info("Server started on port 8000")
		case <-ctx.Done():
			a.log.Info("Context has been canceled before server started")
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
	}

	a.log.Info("Application has successfully stopped")
	return nil
}
