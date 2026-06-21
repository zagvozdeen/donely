package api

import (
	"github.com/labstack/echo/v5"
)

func (a *Application) getMux() *echo.Echo {
	mux := echo.New()
	mux.Logger = a.log.GetSlog()
	mux.Validator = a.createValidator()

	api := mux.Group("/api")

	guest := api.Group("/auth")
	guest.POST("/login", a.login)
	guest.POST("/register", a.register)

	auth := api.Group("/", a.authBearer)
	auth.GET("me", a.getMe)

	return mux
}
