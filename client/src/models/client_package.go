package models

type PackageSent struct {
	Url    string `json:"url"`
	Image  bool   `json:"image"`
	Prompt string `json:"prompt"`
	Video  bool   `json:"vide"`
}
