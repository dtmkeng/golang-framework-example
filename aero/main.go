package main

import (
	"time"

	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()
	app.Use(func(next aero.Handler) aero.Handler {
		return func(ctx aero.Context) error {
			time.Sleep(time.Millisecond * 200)
			next(ctx)
			time.Sleep(time.Millisecond * 200)
			return nil
		}
	})
	configure(app).Run()
}
func configure(app *aero.Application) *aero.Application {
	app.Get("/", func(ctx aero.Context) error {
		return ctx.String("test")
	})
	return app
}
