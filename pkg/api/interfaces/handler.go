package interfaces

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewHandler() *echo.Echo {
	e := echo.New()

	e.GET("/", hello)
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
