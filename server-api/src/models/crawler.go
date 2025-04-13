package models

type Crawler struct {
	Id           int
	PagesToVisit int
	HtmlSelector string
	Status       string
	Query        string
	PagesBodies  []string
}
