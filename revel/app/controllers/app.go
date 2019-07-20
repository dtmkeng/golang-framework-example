package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}
type Stuff struct {
	Foo string
	Bar int
}
func (c App) Index() revel.Result {

	data := make(map[string]interface{})
    stuff := Stuff{Foo: "Test Revel Go", Bar: 999}
    data["stuff"] = stuff
	return c.RenderJSON(data)

}
