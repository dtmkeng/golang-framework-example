package main

import (
	"strconv"

	"github.com/aofei/air"
)

var port = 8080

func airHandler(req *air.Request, res *air.Response) error {
	return res.WriteString(req.Header.Get("name"))
}
func main() {
	a := air.Default
	a.Address = ":" + strconv.Itoa(port)
	a.GET("/hello", airHandler)
	a.Serve()
}
