package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/getCor", func(c *gin.Context) {

		url := c.Query("url")

		fmt.Print("Data: %s", url)
	})

	r.Run("localhost:8080")
}
