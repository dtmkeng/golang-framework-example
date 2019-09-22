package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		responseTime := time.Since(start)

		// Write it to the log
		fmt.Println(responseTime)

		// Make sure to pass the error back!
		// return err
		// next.ServeHTTP(w, r)
	}
}
func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(
	// 		http.StatusOK,
	// 		gin.H{
	// 			"code":  http.StatusOK,
	// 			"error": "Welcome server 02",
	// 		},
	// 	)
	// })
	// r.GET("/welcome", func(c *gin.Context) {
	// 	firstname := c.DefaultQuery("firstname", "Guest")
	// 	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	// })
	r := gin.New()
	r.Use(Logger())
	// r.Use(gin.Logger())
	r.GET("/", home)
	r.Run() // listen and serve on 0.0.0.0:8080
}
func home(c *gin.Context) {
	c.String(200, "Hello")
}
