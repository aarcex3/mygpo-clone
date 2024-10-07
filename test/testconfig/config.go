package testconfig

import (
	"database/sql"
	"log"

	"github.com/aarcex3/mygpo-clone/internals"
	"github.com/gin-gonic/gin"
)

func SetupDatabase() (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	// Create the users table
	if _, err := db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	// Create the tags table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		code TEXT NOT NULL UNIQUE,
		usage INTEGER NOT NULL DEFAULT 0
	);`); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	// Insert data into the tags table
	_, err = db.Exec(`
		INSERT INTO tags (title, code, usage) VALUES 
		('Technology', 'technology', 530),
		('Science', 'science', 410),
		('Health', 'health', 325),
		('Education', 'education', 275),
		('Finance', 'finance', 600),
		('Sports', 'sports', 475),
		('Travel', 'travel', 200),
		('Food', 'food', 150),
		('Art', 'art', 100),
		('History', 'history', 50);
	`)
	if err != nil {
		log.Fatalf("Could not insert data into tags table: %v", err)
	}

	// Return a cleanup function to close the database
	return db, func() {
		db.Close()
	}
}

// Common setup for the tests
func SetupAppWithDB() (*gin.Engine, *sql.DB, func()) {
	app := gin.Default()
	db, cleanup := SetupDatabase()
	internals.SetUpApp(app, db)
	return app, db, cleanup
}
