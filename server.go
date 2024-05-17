package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// func delete_at_index(list []string, index int) []string {
// 	return append(list[:index], list[index+1:]...)
// }

func main() {
	count := 1
	historyCount := 100
	history := []string{}

	fmt.Printf("Wecome to go-gin")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		count++

		if len(history) >= historyCount {
			newStartIndex := len(history) - historyCount - 1
			history = history[newStartIndex:]
		}
		history = append(history, fmt.Sprintf("Serial: %v! Default get endpoint", count))

		ctx.JSON(200, gin.H{
			"message": "You have reached to default get endpoint",
			"history": history,
		})
	})

	r.GET("/:name", func(ctx *gin.Context) {
		count++
		name := ctx.Param("name")

		if len(history) >= historyCount {
			newStartIndex := len(history) - historyCount - 1
			history = history[newStartIndex:]
		}
		history = append(history, fmt.Sprintf("Serial: %v! Named get endpoint: %s", count, name))

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("Hi %s! You have reached to named get endpoint", name),
			"history": history,
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
