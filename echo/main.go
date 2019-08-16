package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echo")
	})
	e.GET("/hello", func(e echo.Context) error {
		return e.String(http.StatusOK, e.QueryParam("name"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
