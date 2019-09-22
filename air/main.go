package main

import (
	"fmt"
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
		start := time.Now()
		err := next(req, res)
		responseTime := time.Since(start)

		// Write it to the log
		fmt.Println(responseTime)

		// Make sure to pass the error back!
		return err
	}
}
func main() {
	a := air.Default
	a.Address = ":" + strconv.Itoa(port)
	a.GET("/", airHandler, midgase)
	a.Serve()
}
