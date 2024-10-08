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
	server *gin.Engine
	db     *sql.DB
	config *config.Config
}

func New(server *gin.Engine, db *sql.DB, config *config.Config) *application {
	return &application{server: server, db: db, config: config}
}

func (a *application) Run() {

	queries := database.New(a.db)

	userRepo := users.Repository(*queries)
	tagRepo := tags.Repository(*queries)

	authService := auth.Service(userRepo)
	tagService := tags.Service(tagRepo)

	authController := auth.Controller(authService)
	directoryContrller := directory.Controller(tagService)

	apiV1 := a.server.Group("/v1")
	auth := apiV1.Group("/auth")
	{
		auth.POST("/registration", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
	}
	directory := apiV1.Group("/")
	{
		directory.GET("/tags", directoryContrller.RetrieveTopTags)
	}

	if err := a.server.Run(fmt.Sprintf("%s:%s", a.config.ServerHost, a.config.ServerPort)); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
