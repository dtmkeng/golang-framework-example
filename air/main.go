package main

import (
	"fmt"
	"github.com/aofei/air"
)

func main() {
	air.Default.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("Hello, 世界")
	})
	fmt.Println("Server start at 8080")
	air.Default.Serve()
}
