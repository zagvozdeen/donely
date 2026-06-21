package api

import (
	"context"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type apiValidator struct {
	validator *validator.Validate
	conform   *mold.Transformer
}

var _ echo.Validator = (*apiValidator)(nil)

func (a *Application) createValidator() *apiValidator {
	return &apiValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
		conform:   modifiers.New(),
	}
}

func (v *apiValidator) Validate(i any) error {
	if err := v.conform.Struct(context.Background(), i); err != nil {
		return echo.ErrInternalServerError.Wrap(err)
	}
	if err := v.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
