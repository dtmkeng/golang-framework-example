package main

import (
	"fmt"
	"time"

	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()
	app.Use(func(next aero.Handler) aero.Handler {
		return func(ctx aero.Context) error {
			// Measure response time
			start := time.Now()
			err := next(ctx)
			responseTime := time.Since(start)

			// Write it to the log
			fmt.Println(responseTime)

			// Make sure to pass the error back!
			return err
		}
	})
	configure(app).Run()
	// app.ListenAndServe(":8080",app)
}

func configure(app *aero.Application) *aero.Application {
	app.Post("/test", func(ctx aero.Context) error {
		return ctx.String("Hello World")
	})
	app.Post("/help", func(ctx aero.Context) error {
		return ctx.String("hello")
	})
	app.Get("/", func(ctx aero.Context) error {
		return ctx.String("test")
	})
	return app
}
