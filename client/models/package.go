package models

type Package struct {
	Url         string `json:"url"`
	Description bool   `json:"description"`
	Image       bool   `json:"image"`
	Video       bool   `json:"vide"`
	Prompt      string `json:"prompt"`
}
