package database

import (
	"log"

	// "github.com/notalim/hopecore-backend/models"
	"github.com/notalim/hopecore-backend/scraper"
)

func SeedQuotes() {
	quotes := scraper.ScrapeQuotes()

	for _, q := range quotes {
		_, err := DB.Exec("INSERT INTO quotes (text, character, source, media_type) VALUES (?, ?, ?, ?)",
			q.Text, q.Character, q.Source, q.MediaType)
		if err != nil {
			log.Printf("Error seeding quote: %v", err)
		}
	}

	log.Printf("Seeded %d quotes", len(quotes))
}