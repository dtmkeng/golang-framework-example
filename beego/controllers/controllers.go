package controllers

import "github.com/astaxie/beego"

// UserController getall
type UserController struct {
	beego.Controller
}
type mystruct struct {
	FieldOne string `json:"field_one"`
}

// Get aLL
func (u *UserController) Get() {
	mystruct2 := &mystruct{
		FieldOne: "keng",
	}
	u.Data["json"] = &mystruct2
	u.ServeJSON()
}
