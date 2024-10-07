package internals

import (
	"database/sql"

	"github.com/aarcex3/mygpo-clone/internals/controllers"
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/aarcex3/mygpo-clone/internals/repositories"
	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetUpApp(app *gin.Engine, db *sql.DB) {

	// Create a new instance of queries
	queries := database.New(db)

	// Setup repos
	userRepo := repositories.NewUserRepository(*queries)
	tagRepo := repositories.NewTagRepository(*queries)

	// Setup services
	authService := services.NewAuthService(userRepo)
	tagService := services.NewTagService(tagRepo)

	// Setup controllers
	authController := controllers.NewAuthController(authService)
	directoryContrller := controllers.NewDirectoryController(tagService)

	// Setup routes
	apiV1 := app.Group("/v1")
	auth := apiV1.Group("/auth")
	{
		auth.POST("/registration", authController.Registration)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
	}
	directory := apiV1.Group("/")
	{
		directory.GET("/tags", directoryContrller.RetrieveTopTags)
	}

}
