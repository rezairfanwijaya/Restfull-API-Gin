package handler

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to index page",
	})
}
