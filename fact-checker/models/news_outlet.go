package models

type NewsOutlet struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	QueryUrl     string `json:"queryUrl"`
	HtmlSelector string `json:"htmlSelector"`
	Language     string `json:"language"`
	Credibility  int    `json:"credibility"`
}
