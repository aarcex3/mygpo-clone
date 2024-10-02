package internals

import (
	"github.com/aarcex3/mygpo-clone/internals/controllers"
	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
)

func SetUpApp(app *gin.Engine) {
	// Setup dependecies
	authService := &services.AuthService{}
	authController := &controllers.AuthController{AuthService: authService}

	apiV1 := app.Group("/v1")

	auth := apiV1.Group("/auth")
	{
		auth.POST("/registration", authController.Registration)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
	}

}
