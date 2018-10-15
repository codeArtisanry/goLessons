package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		router.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong "+fmt.Sprint(time.Now().Unix())+fmt.Sprint(time.Now().String()))
		})
		router.GET("/api/:name", func(c *gin.Context) {
			name := c.Param("name")
			fmt.Println("Hello %s", name)
		})
	}
	router.Run(":8080")
}
