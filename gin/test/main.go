package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	{
		router.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			fmt.Println("Hello %s", name)

		})
	}
	router.Run(":8080")
}
