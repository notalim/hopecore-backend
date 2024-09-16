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