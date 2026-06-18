package api

import "github.com/labstack/echo/v5"

func (a *Application) Validate(i any) error {
	if err := a.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
