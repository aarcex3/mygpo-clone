package testconfig

import (
	"database/sql"
	"log"
)

func SetupDatabase() (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	// Create the user table
	if _, err := db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	// Return a cleanup function to close the database
	return db, func() {
		db.Close()
	}
}
