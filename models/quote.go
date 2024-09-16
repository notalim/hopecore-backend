package models

type Quote struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Character string `json:"character"`
	Anime     string `json:"anime"`
	Source    string `json:"source"`
	MediaType string `json:"media_type"`
}

type Preferences struct {
	ID               int    `json:"id"`
	UserID           string `json:"user_id"`
	Characters       string `json:"characters"`
	UpdateFrequency  int    `json:"update_frequency"`
}
