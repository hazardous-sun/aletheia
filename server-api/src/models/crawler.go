package models

type Crawler struct {
	Id           int
	PagesToVisit int
	HtmlSelector string
	Status       string
	Query        string
	QueryUrl     string
	PagesBodies  []string
}
