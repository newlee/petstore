package server

import (
	"github.com/labstack/echo"
	"net/http"

)

type Product struct {

}

var homeHandler = func(c echo.Context) error {
	return c.String(http.StatusOK, "Pets store")
}

var productListHandler = func(c echo.Context) error {
	products := make([]Product, 0)
	return c.JSON(http.StatusOK, products)
}

func Router(e *echo.Echo){
	e.GET("/", homeHandler)
	e.GET("/products", productListHandler)
}
