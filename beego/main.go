package main

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	bc "github.com/astaxie/beego/context"
)

var port = 8080

func main() {

	app := beego.NewControllerRegister()
	app.Get("/hello", func(ctx *bc.Context) {
		ctx.WriteString("test")
	})
	// beego.Router("/", &controllers.UserController{})
	// beego.Run()

	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: app,
	// }
	// fmt.Println("server start at..", strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), app)
	// s.ListenAndServe()

}
