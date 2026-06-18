package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (a *Application) getMux() http.Handler {
	mux := echo.New()
	mux.Logger = a.log.GetSlog()
	mux.Validator = a

	api := mux.Group("/api")

	guest := api.Group("/auth")
	guest.POST("/login", a.login)
	guest.POST("/register", a.register)

	auth := api.Group("/", a.authBearer)
	auth.GET("me", func(c *echo.Context) error {
		u := c.Get("user")
		return c.JSON(http.StatusOK, u)
	})

	return mux
}
