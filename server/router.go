package server

import (
	"github.com/labstack/echo"
	"net/http"
)

var petRepo PetRepo
var orderRepo OrderRepo

var homeHandler = func(c echo.Context) error {
	return c.String(http.StatusOK, "Pets store")
}

var petListHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, petRepo.All())
}

var petCreateHandler = func(c echo.Context) error {
	pet := &Pet{}
	c.Bind(pet)
	petRepo.Create(*pet)
	return c.JSON(http.StatusCreated, "")
}

var orderCreateHandler = func(c echo.Context) error {
	order := &Order{}
	c.Bind(order)
	orderRepo.Create(*order)
	return c.JSON(http.StatusCreated, "")
}

var orderGetHandler = func(c echo.Context) error {
	telephone := c.Param("telephone")
	order := orderRepo.Find(telephone)
	if order == nil {
		return c.String(http.StatusNotFound, "")
	}
	return c.JSON(http.StatusOK, *order)
}

func Router(e *echo.Echo) {
	petRepo = NewInMemoryPetRepo()
	orderRepo = NewInMemoryOrderRepo()
	e.GET("/", homeHandler)
	e.GET("/pets", petListHandler)
	e.POST("/pets", petCreateHandler)
	e.GET("/orders/:telephone", orderGetHandler)
	e.POST("/orders", orderCreateHandler)
}
