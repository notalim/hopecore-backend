package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/notalim/hopecore-backend/database"
	"github.com/notalim/hopecore-backend/models"
)

func GetQuotes() ([]models.Quote, error) {
	rows, err := database.DB.Query("SELECT id, text, character, anime FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var q models.Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.Character, &q.Anime); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	if len(quotes) == 0 {
		return fetchQuotesFromAPI()
	}

	return quotes, nil
}

func fetchQuotesFromAPI() ([]models.Quote, error) {
	resp, err := http.Get("https://animechan.vercel.app/api/quotes")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quotes []models.Quote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return nil, err
	}

	for _, q := range quotes {
		_, err := database.DB.Exec("INSERT INTO quotes (text, character, anime) VALUES (?, ?, ?)", q.Text, q.Character, q.Anime)
		if err != nil {
			fmt.Printf("Error inserting quote: %v\n", err)
		}
	}

	return quotes, nil
}

func SavePreferences(userID string, characters string, updateFrequency int) error {
	_, err := database.DB.Exec("INSERT OR REPLACE INTO preferences (user_id, characters, update_frequency) VALUES (?, ?, ?)",
		userID, characters, updateFrequency)
	return err
}

func GetPreferences(userID string) (models.Preferences, error) {
	var pref models.Preferences
	err := database.DB.QueryRow("SELECT id, user_id, characters, update_frequency FROM preferences WHERE user_id = ?", userID).
		Scan(&pref.ID, &pref.UserID, &pref.Characters, &pref.UpdateFrequency)
	return pref, err
}
