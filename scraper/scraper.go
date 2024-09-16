package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/notalim/hopecore-backend/models"
)

var shows = []models.Show{
	{IMDbID: "tt2560140", Name: "Attack on Titan", MediaType: "TV Series"},
	// {IMDbID: "tt0944947", Name: "Game of Thrones", MediaType: "TV Series"},
	// {IMDbID: "tt0903747", Name: "Breaking Bad", MediaType: "TV Series"},
	// Add more shows here
}

func ScrapeQuotes() []models.Quote {
	var quotes []models.Quote
	c := colly.NewCollector(
		colly.AllowedDomains("www.imdb.com"),
	)

	c.OnHTML(".ipc-page-content-container", func(e *colly.HTMLElement) {
		show := getShowFromURL(e.Request.URL.Path)
		if show == nil {
			return
		}

		e.ForEach(".ipc-html-content-inner-div", func(_ int, el *colly.HTMLElement) {
			el.ForEach("li", func(_ int, li *colly.HTMLElement) {
				character := strings.TrimSpace(li.ChildText("a.ipc-md-link"))
				quoteText := strings.TrimSpace(li.Text)
				
				// Remove the character name and colon from the quote text
				quoteText = strings.TrimPrefix(quoteText, character)
				quoteText = strings.TrimPrefix(quoteText, ": ")
				quoteText = strings.TrimSpace(quoteText)

				quote := models.Quote{
					Text:      quoteText,
					Character: character,
					Source:    show.Name,
					MediaType: show.MediaType,
				}

				quotes = append(quotes, quote)
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for _, show := range shows {
		url := fmt.Sprintf("https://www.imdb.com/title/%s/quotes/", show.IMDbID)
		c.Visit(url)
	}

	return quotes
}

func getShowFromURL(url string) *models.Show {
	for _, show := range shows {
		if strings.Contains(url, show.IMDbID) {
			return &show
		}
	}
	return nil
}