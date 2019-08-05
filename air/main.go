package main

import (
	"strconv"

	"github.com/aofei/air"
)

var port = 8081

func airHandler(req *air.Request, res *air.Response) error {
	return res.WriteString("Hello, 世界")
}
func main() {
	a := air.New()
	a.Address = ":" + strconv.Itoa(port)
	a.GET("/", airHandler)
	a.Serve()
}
