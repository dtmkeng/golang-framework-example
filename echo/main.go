package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echo")
	})
	e.GET("/hello", func(e echo.Context) error {
		return e.String(http.StatusOK, e.QueryParam("name"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
