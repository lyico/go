package main

import (
	"fmt"
	"github.com/lyico/server/internal/router"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":7777")

	if err != nil {
		return
	}

	fmt.Println("Press Enter key to exit...")
}
