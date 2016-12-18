package main

import (
	"github.com/labstack/echo"
	"github.com/newlee/petstore/server"
)

func main() {
	e := echo.New()
	server.Router(e);
	e.Logger.Fatal(e.Start(":1323"))
}
