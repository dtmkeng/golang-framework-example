package main

import (
	"strconv"
	"time"

	"github.com/aofei/air"
)

var port = 8080

func airHandler(req *air.Request, res *air.Response) error {
	return res.WriteString("Hello, 世界")
}
func midgase(next air.Handler) air.Handler {
	return func(req *air.Request, res *air.Response) error {
		time.Sleep(time.Millisecond * 200)
		next(req, res)
		time.Sleep(time.Millisecond * 200)
		return nil
	}
}
func main() {
	a := air.Default
	a.Address = ":" + strconv.Itoa(port)
	a.GET("/", airHandler, midgase)
	a.Serve()
}
