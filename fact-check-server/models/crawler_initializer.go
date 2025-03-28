package models

type CrawlerInitializer struct {
	Query        string `json:"query"`
	PagesToVisit int    `json:"pagesToVisit"`
}
