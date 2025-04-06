package models

type PackageReceived struct {
	Image  bool   `json:"image"`
	Video  bool   `json:"video"`
	Prompt string `json:"prompt"`
	Url    string `json:"url"`
}
