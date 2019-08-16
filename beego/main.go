package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	bc "github.com/astaxie/beego/context"
)

var port = 8080

func main() {

	app := beego.NewControllerRegister()
	app.Get("/", func(ctx *bc.Context) {
		ctx.WriteString("Hello World")
	})
	app.Get("/hello", func(ctx *bc.Context) {
		ctx.WriteString(ctx.Input.Query("name"))
	})
	// beego.Router("/", &controllers.UserController{})
	// beego.Run()
	fmt.Println("server start at..", strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), app)

}
