package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/zagvozdeen/donely/internal/store/models"
	"golang.org/x/crypto/bcrypt"
)

type authRequest struct {
	Email    string `json:"email" mod:"trim,lcase" validate:"required,email,min=5,max=100"`
	Password string `json:"password" mod:"trim" validate:"required,min=8,max=100"`
}

func (a *Application) login(c *echo.Context) error {
	req := new(authRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	user, err := a.store.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return echo.ErrNotFound.Wrap(err)
		}
		return echo.ErrInternalServerError.Wrap(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return echo.ErrUnauthorized.Wrap(err)
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        user.UUID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365)),
	})
	token, err := t.SignedString([]byte(a.cfg.AppSecret))
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

type registerRequest struct {
	FirstName            string `json:"first_name" mod:"trim" validate:"required,max=255"`
	LastName             string `json:"last_name" mod:"trim" validate:"required,max=255"`
	Email                string `json:"email" mod:"trim,lcase" validate:"required,email,max=256"`
	Password             string `json:"password" mod:"trim" validate:"required,min=8,max=72"`
	PasswordConfirmation string `json:"password_confirmation" mod:"trim" validate:"required,eqfield=Password"`
}

func (a *Application) register(c *echo.Context) error {
	req := new(registerRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	_, err := a.store.GetUserByEmail(c.Request().Context(), req.Email)
	if err == nil {
		return echo.ErrBadRequest.Wrap(fmt.Errorf("user already exists"))
	}
	if !errors.Is(err, models.ErrUserNotFound) {
		return echo.ErrInternalServerError.Wrap(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}
	uid, err := uuid.NewV7()
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}
	user := &models.User{
		UUID:      uid,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = a.store.CreateUser(c.Request().Context(), user)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (a *Application) authBearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return echo.ErrUnauthorized.Wrap(errors.New("missing authorization header"))
		}
		if !strings.HasPrefix(header, "Bearer ") {
			return echo.ErrUnauthorized.Wrap(errors.New("invalid authorization header"))
		}
		token := strings.TrimPrefix(header, "Bearer ")
		var claims jwt.RegisteredClaims
		t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
			return []byte(a.cfg.AppSecret), nil
		})
		if err != nil {
			return echo.ErrUnauthorized.Wrap(fmt.Errorf("failed to parse token: %w", err))
		}
		if !t.Valid {
			return echo.ErrUnauthorized.Wrap(errors.New("invalid token"))
		}
		var uid uuid.UUID
		uid, err = uuid.Parse(claims.ID)
		if err != nil {
			return echo.ErrUnauthorized.Wrap(fmt.Errorf("invalid token claims: %w", err))
		}
		var user *models.User
		user, err = a.store.GetUserByUUID(c.Request().Context(), uid)
		if err != nil {
			if errors.Is(err, models.ErrUserNotFound) {
				return echo.ErrUnauthorized.Wrap(fmt.Errorf("user not found: %w", err))
			}
			return echo.ErrInternalServerError.Wrap(fmt.Errorf("failed to get user: %w", err))
		}
		c.Set("user", user)
		return next(c)
	}
}

func (a *Application) getMe(c *echo.Context) error {
	user, ok := c.Get("user").(*models.User)
	if !ok {
		return echo.ErrUnauthorized.Wrap(errors.New("user not authenticated"))
	}
	return c.JSON(http.StatusOK, user)
}
