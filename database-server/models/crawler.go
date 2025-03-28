package models

type Crawler struct {
	Id           int
	PagesToVisit int
	Query        string
	QueryUrl     string
	HtmlSelector string
	Status       string
	PagesBodies  []string
}
