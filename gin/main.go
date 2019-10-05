package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func echoMiddlew() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(time.Millisecond * 200)
		c.Next()
		time.Sleep(time.Millisecond * 200)
	}
}
func main() {
	r := gin.New()
	r.Use(echoMiddlew())
	r.GET("/", home)
	r.Run() // listen and serve on 0.0.0.0:8080
}
func home(c *gin.Context) {
	c.String(200, "Hello")
}
