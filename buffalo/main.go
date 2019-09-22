package main

import (
	"fmt"
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

func MyMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// do some work before calling the next handler
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

	app := buffalo.New(buffalo.Options{
		Env:          ENV,
		SessionStore: sessions.Null{},
		PreWares: []buffalo.PreWare{
			cors.Default().Handler,
		},
		SessionName: "_coke_session",
		Addr:        ":8080",
	})
	// app.Use(buffalo.RequestLoggerFunc)
	// app.Use(contenttype.Set("application/json"))
	app.Use(MyMiddleware)
	app.GET("/", HomeHandler)
	// app.Env = "production"
	// app.GET("/hello", func(c buffalo.Context) error {
	// 	return c.Render(200, r.String(c.Params())
	// })
	// app.Options.Addr := 8080
	app.Serve()

}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
