package test

import (
	"database/sql"
	"log"

	"github.com/aarcex3/mygpo-clone/config"
)

func SetupTestDatabase(config *config.Config) (*sql.DB, func()) {
	db, err := sql.Open(config.DatabaseEngine, config.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		code TEXT NOT NULL UNIQUE,
		usage INTEGER NOT NULL DEFAULT 0
	);`); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

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

	return db, func() {
		db.Close()
	}
}
