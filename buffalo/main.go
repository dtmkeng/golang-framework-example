package main

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

var r *render.Engine

// ENV ...
var ENV = envy.Get("GO_ENV", "production")

// MyMiddleware ...
func MyMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		time.Sleep(time.Millisecond * 200)
		next(c)
		time.Sleep(time.Millisecond * 200)
		return nil
	}
}
func main() {

	app := buffalo.New(buffalo.Options{
		Env:          ENV,
		SessionStore: sessions.Null{},
		PreWares: []buffalo.PreWare{
			cors.Default().Handler,
		},
		SessionName: "_coke_session",
		Addr:        ":8080",
	})
	app.Use(MyMiddleware)
	app.GET("/", HomeHandler)
	app.Serve()

}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
