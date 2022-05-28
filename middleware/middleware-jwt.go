package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func MiddlewareJWT(c *gin.Context) {
	// ambil cookie
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	// ambil token
	jwtToken, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123"), nil
	})
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Token Tidak Ditemukan",
		})
		c.Abort()
		return
	}

	// cek token
	if !jwtToken.Valid {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Token Tidak Valid",
		})
		c.Abort()
		return
	}

	// jika berhasil maka lanjutkan ke handler
	c.Next()
}
