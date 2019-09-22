package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		start := time.Now()
		err := next(c)
		responseTime := time.Since(start)

		// Write it to the log
		fmt.Println(responseTime)

		// Make sure to pass the error back!
		return err
	}
}
func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(process)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echo")
	})
	e.GET("/hello", func(e echo.Context) error {
		return e.String(http.StatusOK, e.QueryParam("name"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
