package main

import (
	"log"
	"restfull-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// insiasi gin router
	route := gin.Default()

	// definisi router dan route handler (di import dari handler/handler.go)
	route.GET("/", handler.IndexHandler)

	// run server
	log.Println("Server berjalan pada http://localhost:8080")
	route.Run(":8080")

}
