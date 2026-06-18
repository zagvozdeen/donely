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
	Email    string `json:"email" validate:"required,email,min=5,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
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
	//req, res := validation.Validate[registerRequest](r.Context(), r.Body, a.conform, a.validate)
	//if res != nil {
	//	return res
	//}
	_, err := a.store.GetUserByEmail(c.Request().Context(), req.Email)
	if err == nil {
		return echo.ErrBadRequest.Wrap(fmt.Errorf("user already exists"))
		//return core.Err(http.StatusConflict, fmt.Errorf("user already exists"))
	}
	if !errors.Is(err, models.ErrUserNotFound) {
		return echo.ErrInternalServerError.Wrap(err)
		//return core.Err(http.StatusInternalServerError, err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
		//return core.Err(http.StatusInternalServerError, fmt.Errorf("failed to hash password: %w", err))
	}
	uid, err := uuid.NewV7()
	if err != nil {
		return echo.ErrInternalServerError.Wrap(err)
		//return core.Err(http.StatusInternalServerError, fmt.Errorf("failed to generate uuid v7: %w", err))
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
		//return core.Err(http.StatusInternalServerError, fmt.Errorf("failed to create user: %w", err))
	}
	return c.JSON(http.StatusCreated, user)
}

func (a *Application) authBearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return echo.ErrUnauthorized.Wrap(errors.New("missing authorization header"))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("missing authorization header"))
		}
		if !strings.HasPrefix(header, "Bearer ") {
			return echo.ErrUnauthorized.Wrap(errors.New("invalid authorization header"))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
		}
		token := strings.TrimPrefix(header, "Bearer ")
		var claims jwt.RegisteredClaims
		t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
			return []byte(a.cfg.AppSecret), nil
		})
		if err != nil {
			return echo.ErrUnauthorized.Wrap(fmt.Errorf("failed to parse token: %w", err))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to parse token: %w", err))
		}
		if !t.Valid {
			return echo.ErrUnauthorized.Wrap(errors.New("invalid token"))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token"))
		}
		var uid uuid.UUID
		uid, err = uuid.Parse(claims.ID)
		if err != nil {
			return echo.ErrUnauthorized.Wrap(fmt.Errorf("invalid token claims: %w", err))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %w", err))
		}
		var user *models.User
		user, err = a.store.GetUserByUUID(c.Request().Context(), uid)
		if err != nil {
			if errors.Is(err, models.ErrUserNotFound) {
				return echo.ErrUnauthorized.Wrap(fmt.Errorf("user not found: %w", err))
				//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("user not found: %w", err))
			}
			return echo.ErrInternalServerError.Wrap(fmt.Errorf("failed to get user: %w", err))
			//return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to get user: %w", err))
		}
		c.Set("user", user)
		return next(c)
	}
}

//func (a *Application) auth(ctx *core.Context) (*models.User, error) {
//	header := ctx.Request().Header.Get("Authorization")
//	if header == "" {
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("missing authorization header"))
//	}
//	if !strings.HasPrefix(header, "Bearer ") {
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
//	}
//	token := strings.TrimPrefix(header, "Bearer ")
//	var claims jwt.RegisteredClaims
//	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
//		return []byte(a.cfg.AppSecret), nil
//	})
//	if err != nil {
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to parse token: %w", err))
//	}
//	if !t.Valid {
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token"))
//	}
//	var uid uuid.UUID
//	uid, err = uuid.Parse(claims.ID)
//	if err != nil {
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %w", err))
//	}
//	var user *models.User
//	user, err = a.store.GetUserByUUID(ctx, uid)
//	if err != nil {
//		if errors.Is(err, models.ErrUserNotFound) {
//			return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("user not found: %w", err))
//		}
//		return nil, ctx.Error(http.StatusUnauthorized, fmt.Errorf("failed to get user: %w", err))
//	}
//	return user, nil
//}
//
////func (s *Service) guest(fn core.GuestHandlerFunc) http.HandlerFunc {
////	return func(w http.ResponseWriter, r *http.Request) {
////		fn(r).Response(w, s.log)
////	}
////}
////
////func (s *Service) auth(fn core.HandlerFunc) http.HandlerFunc {
////	return func(w http.ResponseWriter, r *http.Request) {
////		req, user, res := s.checkAuth(r, r.Header.Get("Authorization"))
////		if res == nil {
////			res = fn(req, user)
////		}
////		res.Response(w, s.log)
////	}
////}
////
////func (s *Service) role(fn core.HandlerFunc, role enums.UserRole) http.HandlerFunc {
////	return func(w http.ResponseWriter, r *http.Request) {
////		req, user, res := s.checkAuth(r, r.Header.Get("Authorization"))
////		if res == nil {
////			if user.Role.Priority() < role.Priority() {
////				res = core.Err(http.StatusForbidden, fmt.Errorf("forbidden: insufficient permissions"))
////			} else {
////				res = fn(req, user)
////			}
////		}
////		res.Response(w, s.log)
////	}
////}
////
////func (s *Service) checkAuth(r *http.Request, token string) (*http.Request, *models.User, core.Response) {
////	switch {
////	case strings.HasPrefix(token, "tma "):
////		return s.authTMA(r, token)
////	case strings.HasPrefix(token, "Bearer "):
////		return s.authBearer(r, token)
////	default:
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("missing authorization header"))
////	}
////}
////
////func (s *Service) authTMA(r *http.Request, token string) (*http.Request, *models.User, core.Response) {
////	token = strings.TrimPrefix(token, "tma ")
////	values, err := url.ParseQuery(token)
////	if err != nil {
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("failed to parse tma token: %w", err))
////	}
////	u, ok := bot.ValidateWebappRequest(values, s.cfg.Telegram.BotToken)
////	if !ok {
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("invalid tma token"))
////	}
////	var user *models.User
////	user, err = s.store.GetUserByTID(r.Context(), u.ID)
////	if err != nil {
////		if errors.Is(err, models.ErrNotFound) {
////			return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("tma user not found: %w", err))
////		}
////		return nil, nil, core.Err(http.StatusInternalServerError, fmt.Errorf("failed to load user: %w", err))
////	}
////	return r.WithContext(context.WithValue(r.Context(), "source", enums.OrderSourceTMA)), user, nil
////}
////
////func (s *Service) authBearer(r *http.Request, token string) (*http.Request, *models.User, core.Response) {
////	token = strings.TrimPrefix(token, "Bearer ")
////	var claims jwt.RegisteredClaims
////	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
////		return []byte(s.cfg.App.Secret), nil
////	})
////	if err != nil {
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("failed to parse token: %w", err))
////	}
////	if !t.Valid {
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("invalid token"))
////	}
////	id, err := strconv.Atoi(claims.ID)
////	if err != nil {
////		return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("invalid token: %w, id=%s", err, claims.ID))
////	}
////	var user *models.User
////	user, err = s.store.GetUserByID(r.Context(), id)
////	if err != nil {
////		if errors.Is(err, models.ErrNotFound) {
////			return nil, nil, core.Err(http.StatusUnauthorized, fmt.Errorf("user not found: %w", err))
////		}
////		return nil, nil, core.Err(http.StatusInternalServerError, fmt.Errorf("failed to load user: %w", err))
////	}
////	return r.WithContext(context.WithValue(r.Context(), "source", enums.OrderSourceSPA)), user, nil
////}
