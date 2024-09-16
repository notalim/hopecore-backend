package main

import (
	"fmt"
	"log"

	"github.com/notalim/hopecore-backend/database"
	"github.com/notalim/hopecore-backend/scraper"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	// Drop the quotes table
	_, err := database.DB.Exec("DROP TABLE IF EXISTS quotes")
	if err != nil {
		log.Printf("Error dropping table: %v", err)
	}

	// Recreate the quotes table
	_, err = database.DB.Exec(`
		CREATE TABLE quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			character TEXT NOT NULL,
			source TEXT NOT NULL,
			media_type TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Error recreating table:", err)
	}

	quotes := scraper.ScrapeQuotes()
	for _, q := range quotes {
		_, err := database.DB.Exec("INSERT INTO quotes (text, character, source, media_type) VALUES (?, ?, ?, ?)",
			q.Text, q.Character, q.Source, q.MediaType)
		if err != nil {
			log.Printf("Error inserting quote: %v", err)
		}
	}

	fmt.Printf("Scraped and inserted %d quotes\n", len(quotes))

	// Print the quotes for verification
	rows, err := database.DB.Query("SELECT * FROM quotes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var text, character, source, mediaType string
		err := rows.Scan(&id, &text, &character, &source, &mediaType)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s - %s (%s, %s)\n", id, text, character, source, mediaType)
	}
}