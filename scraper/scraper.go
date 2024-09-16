package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/notalim/hopecore-backend/models"
)

func ScrapeQuotes() []models.Quote {
	var quotes []models.Quote
	quoteCount := 0
	maxQuotes := 10

	c := colly.NewCollector(
		colly.AllowedDomains("www.imdb.com"),
	)

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		if quoteCount >= maxQuotes {
			return
		}

		quoteText := strings.TrimSpace(e.ChildText(".quote-text"))
		character := strings.TrimSpace(e.ChildText(".quote-character"))
		source := strings.TrimSpace(e.ChildAttr("a.quote-source", "title"))
		mediaType := determineMediaType(source)

		quote := models.Quote{
			Text:      quoteText,
			Character: character,
			Source:    source,
			MediaType: mediaType,
		}

		quotes = append(quotes, quote)
		quoteCount++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Visit multiple pages to get a variety of quotes
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("https://www.imdb.com/search/title-text/?quotes=on&sort=num_votes,desc&page=%d", i)
		c.Visit(url)
	}

	return quotes
}

func determineMediaType(source string) string {
	lowercaseSource := strings.ToLower(source)
	if strings.Contains(lowercaseSource, "tv series") || strings.Contains(lowercaseSource, "tv mini series") {
		return "TV Show"
	} else if strings.Contains(lowercaseSource, "anime") {
		return "Anime"
	} else {
		return "Movie"
	}
}