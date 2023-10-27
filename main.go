package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {

		id := context.Query("id")

		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
			"id":   id,
		}

		context.JSON(http.StatusOK, data)
	})

	err := r.Run(":7070")

	if err != nil {
		return
	}

	fmt.Println("Press Enter key to exit...")
}
