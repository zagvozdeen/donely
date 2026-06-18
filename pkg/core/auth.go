package core

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zagvozdeen/donely/internal/store/models"
)

func (s *Service) authBearer(ctx *Context) (*models.User, error) {
	header := ctx.Request().Header.Get("Authorization")
	if header == "" {
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("missing authorization header"))
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
	}
	token := strings.TrimPrefix(header, "Bearer ")
	var claims jwt.RegisteredClaims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(s.cfg.AppSecret), nil
	})
	if err != nil {
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to parse token: %w", err))
	}
	if !t.Valid {
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token"))
	}
	var uid uuid.UUID
	uid, err = uuid.Parse(claims.ID)
	if err != nil {
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %w", err))
	}
	var user *models.User
	user, err = s.store.GetUserByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("user not found: %w", err))
		}
		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to get user: %w", err))
	}
	return user, nil
}
