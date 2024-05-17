package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

//	func delete_at_index(list []string, index int) []string {
//		return append(list[:index], list[index+1:]...)
//	}
func reverse_slice(originalSlice []string) []string {

	slice := make([]string, len(originalSlice))
	count := copy(slice, originalSlice)
	fmt.Println("Copied: ", count)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

type Response struct {
	Item1 string   `json:"item1"`
	Item2 []string `json:"item2"`
}

func main() {
	count := 0
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

		reversed_history := reverse_slice(history)

		res := Response{
			Item1: "Howdy partner!? You have reached the default get endpoint. Good day",
			Item2: reversed_history,
		}

		ctx.JSON(200, gin.H{
			"res": res,
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

		reversed_history := reverse_slice(history)

		res := Response{
			Item1: fmt.Sprintf("Good day Sir %v, You have reached the named get endpoint.", name),
			Item2: reversed_history,
		}

		ctx.JSON(200, gin.H{
			"res": res,
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
