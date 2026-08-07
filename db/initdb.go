package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/pantry.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if _, err := db.Exec("DROP TABLE IF EXISTS pantry"); err != nil {
		log.Fatalf("Error dropping table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE pantry (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ingredient TEXT NOT NULL UNIQUE,
			ingredientType TEXT,
			quantity FLOAT NOT NULL,
			units TEXT NOT NULL,
			dateUpdated DATETIME NOT NULL,
			expectedExpiry DATETIME NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println(".db created with pantry table.")
}
