package models

type NewsOutlet struct {
	Id           int    `json:"id"`
	Credibility  int    `json:"credibility"`
	HtmlSelector string `json:"htmlSelector"`
	Language     string `json:"language"`
	Name         string `json:"name"`
	QueryUrl     string `json:"queryUrl"`
}
