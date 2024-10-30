package main

import (
	"go-crud-challenge/delivery/routers"
)

func main() {
	router := routers.SetupRouter()

	// Start the server with error handling
	if err := router.Run(":8000"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
