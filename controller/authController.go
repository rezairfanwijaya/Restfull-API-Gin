package controller

import (
	"fmt"
	"net/http"
	"restfull-api/models"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
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

type LoginInput struct {
	Password string `json:"password" binding:"required" `
	Email    string `json:"email" binding:"required" validate:"email"`
}

var KEY = "1223bhvjerhbguygue"

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
			"code":    400,
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
			"code":    400,
		})

		return
	}

	// encrypt password
	passEnc, err := bcrypt.GenerateFromPassword([]byte(dataInput.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"code":    400,
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
	// inisiasi struct login input
	var dataInput LoginInput

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
			"code":    400,
		})

		return
	}

	// cek email
	var sisfo models.Sisfo
	models.DB.Where("email = ?", dataInput.Email).First(&sisfo)
	if sisfo.ID == 0 {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Email tidak terdaftar",
			"code":    400,
		})

		return
	}

	// cek password
	err = bcrypt.CompareHashAndPassword(sisfo.Password, []byte(dataInput.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Password salah",
			"code":    400,
		})

		return
	}

	// buatkan claims
	exp := time.Now().Add(time.Hour * 24)
	id := sisfo.ID
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: &jwt.Time{exp},
		Issuer:    string(id),
	})

	// tanda tangani claim
	tokenJWT, err := claims.SignedString([]byte(KEY))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"code":    400,
		})

		return
	}

	// masukan token ke cookie
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   tokenJWT,
		Expires: exp,
	}

	// simpan cookie
	http.SetCookie(c.Writer, &cookie)

	c.JSON(200, gin.H{
		"status": "success login",
		"code":   200,
	})

}
