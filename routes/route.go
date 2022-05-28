package routes

import (
	"restfull-api/controller"
	"restfull-api/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(App *gin.Engine) {

	// route yang tidak memerlukan auth (yang membuat auth)
	App.POST("/register", controller.Register)
	App.POST("/login", controller.Login)

	// membuat route group untuk auth (yang memakai auth)
	Auth := App.Group("/")
	Auth.Use(middleware.MiddlewareJWT)
	{
		Auth.GET("/", controller.Home)
		Auth.GET("/mahasiswa", controller.TampilData)
		Auth.POST("/mahasiswa", controller.TambahMahasiswa)
		Auth.PUT("/mahasiswa/:nim", controller.EditMahasiswa)
		Auth.DELETE("/mahasiswa/:nim", controller.HapusMahasiswa)
	}
}
