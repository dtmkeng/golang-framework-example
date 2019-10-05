package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		time.Sleep(time.Millisecond * 200)
		next(c)
		time.Sleep(time.Millisecond * 200)
		return nil
	}
}
func main() {
	e := echo.New()
	e.Use(process)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echo")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
