package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/auth"
	"github.com/aarcex3/mygpo-clone/internals/database"
	"github.com/aarcex3/mygpo-clone/internals/directory"
	"github.com/aarcex3/mygpo-clone/internals/tags"
	"github.com/aarcex3/mygpo-clone/internals/users"

	"github.com/gin-gonic/gin"
)

type application struct {
	router *gin.Engine
	db     *sql.DB
	config *config.Config
}

func New(router *gin.Engine, db *sql.DB, config *config.Config) *application {
	app := &application{router: router, db: db, config: config}
	queries := database.New(app.db)

	userRepo := users.Repository(*queries)
	tagRepo := tags.Repository(*queries)

	authService := auth.Service(userRepo, app.config)
	tagService := tags.Service(tagRepo)

	authController := auth.Controller(authService)
	directoryController := directory.Controller(tagService)

	apiV1 := app.router.Group("/v1")
	auth := apiV1.Group("/auth")
	{
		auth.POST("/registration", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
	}
	directory := apiV1.Group("/")
	{
		directory.GET("/tags", directoryController.RetrieveTopTags)
	}

	return app

}

func (a *application) Run() {
	if err := a.router.Run(fmt.Sprintf("%s:%s", a.config.ServerHost, a.config.ServerPort)); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
