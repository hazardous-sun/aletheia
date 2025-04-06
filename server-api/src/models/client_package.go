package models

type PackageReceived struct {
	Url    string `json:"url"`
	Image  bool   `json:"image"`
	Video  bool   `json:"video"`
	Prompt string `json:"prompt"`
}
