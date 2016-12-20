package server

import (
	"github.com/labstack/echo"
	"net/http"
)

var repo ProductRepo = NewInMemoryProductRepo()

var homeHandler = func(c echo.Context) error {
	return c.String(http.StatusOK, "Pets store")
}

var productListHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, repo.All())
}

var productCreateHandler = func(c echo.Context) (err error) {
	product := &Product{}
	c.Bind(product)
	repo.Create(*product)
	return c.JSON(http.StatusCreated, "")
}

func Router(e *echo.Echo){
	e.GET("/", homeHandler)
	e.GET("/products", productListHandler)
	e.POST("/products", productCreateHandler)
}
