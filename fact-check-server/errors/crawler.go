package custom_errors

const (
	CrawlerEmptyQuery        = "crawler query cannot be empty"
	CrawlerEmptyQueryUrl     = "crawler query url cannot be empty"
	CrawlerEmptyHtmlSelector = "crawler html selector cannot be empty"
	CrawlerFilledPagesBodies = "crawler filled pages bodies needs to be empty"
	CrawlerClosingPageError  = "crawler did not close the page properly"
)
