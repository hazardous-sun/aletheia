package models

type NewsOutlet struct {
	Id           int    `json:"id"`
	QueryUrl     string `json:"queryUrl"`
	HtmlSelector string `json:"htmlSelector"`
	Name         string `json:"name"`
	Language     string `json:"language"`
	Credibility  int    `json:"credibility"`
}
