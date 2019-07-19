package main

import (
	"github.com/astaxie/beego"
	"github.com/dtmkeng/framework-exmaple/beego/controllers"
)

func main() {

	beego.Router("/", &controllers.UserController{})
	beego.Run()

}
