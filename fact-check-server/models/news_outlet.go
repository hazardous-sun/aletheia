package models

type NewsOutlet struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Language    string `json:"language"`
	Credibility int    `json:"credibility"`
}
