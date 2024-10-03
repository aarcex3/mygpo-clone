package internals

import (
	"database/sql"
	"log"

	"github.com/aarcex3/mygpo-clone/internals/controllers"
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
)

func SetUpApp(app *gin.Engine) {
	// Setup db
	db, err := sql.Open("sqlite3", "./example.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	queries := database.New(db)

	// Setup repos
	userRepo := repositories.NewUserRepository(*queries)

	// Setup services
	authService := services.NewAuthService(userRepo)

	//Setup controllers
	authController := controllers.NewAuthController(authService)

	apiV1 := app.Group("/v1")

	auth := apiV1.Group("/auth")
	{
		auth.POST("/registration", authController.Registration)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
	}

}
