package main

import (
	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()
	app.Get("/hello", aeroHandler)
	app.Run()
}
func aeroHandler(a aero.Context) error {
	return a.String("test")
}
