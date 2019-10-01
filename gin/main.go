package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/hello", ginHandler)
	r.Run()
}
func ginHandler(c *gin.Context) {
	c.String(200, "test")
}
