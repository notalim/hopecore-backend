package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./hopecore.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createQuotesTable := `
	CREATE TABLE IF NOT EXISTS quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		character TEXT NOT NULL,
		source TEXT NOT NULL,
		media_type TEXT NOT NULL
	);`

	_, err := DB.Exec(createQuotesTable)
	if err != nil {
		log.Fatal(err)
	}
}