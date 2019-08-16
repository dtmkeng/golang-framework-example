package main

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	contenttype "github.com/gobuffalo/mw-contenttype"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

var r *render.Engine

// ENV ...
var ENV = envy.Get("GO_ENV", "development")

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
	app.Use(paramlogger.ParameterLogger)
	app.Use(contenttype.Set("application/json"))

	app.GET("/", HomeHandler)
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
