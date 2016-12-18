package server

import (
	"github.com/labstack/echo"
	"net/http"
)
func Router(e *echo.Echo){
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pets store")
	})
}
