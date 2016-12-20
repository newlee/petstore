package server

import (
	"github.com/labstack/echo"
	"net/http"

)

type Product struct {
	Name string `json:"name" form:"name" query:"name"`
}
var products = make([]Product, 0)

var homeHandler = func(c echo.Context) error {
	return c.String(http.StatusOK, "Pets store")
}

var productListHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

var productCreateHandler = func(c echo.Context) (err error) {
	p := new(Product)
	c.Bind(p)
	products = append(products,*p)
	return c.JSON(http.StatusCreated, "")
}

func Router(e *echo.Echo){
	e.GET("/", homeHandler)
	e.GET("/products", productListHandler)
	e.POST("/products", productCreateHandler)
}
