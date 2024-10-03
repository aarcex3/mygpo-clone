package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aarcex3/mygpo-clone/internals"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin engine
	app := gin.Default()
	// Setup db
	db, err := sql.Open("sqlite3", "./mygpo-clone.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	// Setup the application and get the db connection
	internals.SetUpApp(app, db)

	// Start the server in a separate goroutine
	go func() {
		if err := app.Run(":8000"); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Wait for a termination signal to gracefully shut down the application
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Close the database connection
	if err := db.Close(); err != nil {
		log.Fatal("Failed to close database:", err)
	}

	log.Println("Server gracefully stopped.")
}
