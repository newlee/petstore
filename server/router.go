package server

import (
	"github.com/labstack/echo"
	"net/http"
)

var repo PetRepo = NewInMemoryPetRepo()

var homeHandler = func(c echo.Context) error {
	return c.String(http.StatusOK, "Pets store")
}

var petListHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, repo.All())
}

var petCreateHandler = func(c echo.Context) (err error) {
	pet := &Pet{}
	c.Bind(pet)
	repo.Create(*pet)
	return c.JSON(http.StatusCreated, "")
}

func Router(e *echo.Echo) {
	e.GET("/", homeHandler)
	e.GET("/pets", petListHandler)
	e.POST("/pets", petCreateHandler)
}
