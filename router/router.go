package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ohunyk/webserver/interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo{
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/users", func(ctx echo.Context) error {
		return c.GetUsers(ctx)
	})
	return e
}
