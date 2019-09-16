package main

import (
	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()
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
	return app
}
