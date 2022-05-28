package controller

import (
	"fmt"
	"restfull-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// struct input user register
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" `
	Email    string `json:"email" binding:"required" validate:"email"`
}

func Register(c *gin.Context) {
	// inisiasi register input
	var dataInput RegisterInput

	// binding data input
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		// bikin variabel penampung error
		var error []string

		// ambil error jika ada
		for _, err := range err.(validator.ValidationErrors) {
			errMSG := fmt.Sprintf("Error on filed: %s, condition: %s", err.Field(), err.ActualTag())
			error = append(error, errMSG)
		}

		c.JSON(400, gin.H{
			"status":  "error",
			"message": error,
		})

		return
	}

	// cek apakah ada email yang sama
	var sisfo models.Sisfo
	models.DB.Where("email = ?", dataInput.Email).First(&sisfo)
	if sisfo.ID != 0 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Email sudah terdaftar",
		})

		return
	}

	// encrypt password
	passEnc, err := bcrypt.GenerateFromPassword([]byte(dataInput.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	// input data ke struct sisfo
	newSisfo := models.Sisfo{
		Username: dataInput.Username,
		Password: passEnc,
		Email:    dataInput.Email,
	}

	// save ke db
	models.DB.Create(&newSisfo)

	c.JSON(200, gin.H{
		"status": "success",
		"data":   newSisfo,
		"code":   200,
	})
}

func Login(c *gin.Context) {
}
