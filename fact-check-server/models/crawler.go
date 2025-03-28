package models

type Crawler struct {
	Id           int
	PagesToVisit int
	Query        string
	QueryUrl     string
	HtmlSelector string
	PagesBodies  []string
}
