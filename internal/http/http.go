package http

import (
	"fmt"
	"net/http"

	"projectONE/internal/app/users"
	"projectONE/internal/config"
	"projectONE/internal/factory"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = config.Get().Server.App
		VERSION = config.Get().Server.Version
		// HOST    = config.Get().Server.Host
		// SCHEME  = config.Get().Server.Scheme
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// users
	users.NewHandler(f).Route(e.Group("/users"))

}
