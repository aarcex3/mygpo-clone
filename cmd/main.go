package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aarcex3/mygpo-clone/config"
	"github.com/aarcex3/mygpo-clone/internals/app"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config := config.LoadConfig("dev")
	server := gin.Default()
	db, err := sql.Open(config.DatabaseEngine, config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(server, db, config)
	app.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := db.Close(); err != nil {
		log.Fatal("Failed to close database:", err)
	}

	log.Println("Server gracefully stopped.")
}
