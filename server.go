package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Wecome to go-gin")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "You have reached to default get endpoint",
		})
	})

	r.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("Hi %s! You have reached to named get endpoint", name),
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
