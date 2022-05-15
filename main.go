package main

import (
	"log"
	"restfull-api/handler"
	"restfull-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// import koneksi database
	db := models.Connect()

	// insiasi gin router
	route := gin.Default()

	// pasang koneksi database ke router
	// code ini akan dijalankan pertama sebelum masuk ke handler
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next() // jika koneksi sudah berhasil maka lanjutkan ke handler
	})

	// definisi router dan route handler (di import dari handler/handler.go)
	route.GET("/", handler.IndexHandler)

	// run server
	log.Println("Server berjalan pada http://localhost:8080")
	route.Run(":8080")

}
