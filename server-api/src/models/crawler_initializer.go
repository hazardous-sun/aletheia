package models

type CrawlerInitializer struct {
	PagesToVisit int    `json:"pagesToVisit"`
	Query        string `json:"query"`
}
