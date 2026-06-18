package core

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"net/http"
	"time"

	"github.com/zagvozdeen/donely/internal/config"
	"github.com/zagvozdeen/donely/internal/logger"
	"github.com/zagvozdeen/donely/internal/store"
	"github.com/zagvozdeen/donely/internal/store/models"
)

type Service struct {
	cfg   *config.Config
	log   *logger.Logger
	store *store.Store
	auth  func(*Context) (*models.User, error)
}

func NewService(cfg *config.Config, log *logger.Logger, store *store.Store, auth func(*Context) (*models.User, error)) *Service {
	return &Service{
		cfg:   cfg,
		log:   log,
		store: store,
		auth:  auth,
	}
}

type Context struct {
	u *models.User
	w http.ResponseWriter
	r *http.Request
}

var _ context.Context = (*Context)(nil)

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.r.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.r.Context().Done()
}

func (c *Context) Err() error {
	return c.r.Context().Err()
}

func (c *Context) Value(key any) any {
	return c.r.Context().Value(key)
}

func (c *Context) User() *models.User {
	return c.u
}

func (c *Context) Request() *http.Request {
	return c.r
}

func (c *Context) JSON(statusCode int, value any) error {
	c.w.WriteHeader(statusCode)
	err := json.MarshalWrite(c.w, value)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	return nil
}

func (c *Context) Error(statusCode int, err error) error {
	http.Error(c.w, err.Error(), statusCode)
	return nil
}
