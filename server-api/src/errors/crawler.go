package server_errors

const (
	CrawlerReady     = "crawler is ready"
	CrawlerRunning   = "crawler is running"
	CrawlerSucceeded = "crawler successfully crawled"
)

const (
	CrawlerEmptyQuery        = "crawler query cannot be empty"
	CrawlerEmptyQueryUrl     = "crawler query url cannot be empty"
	CrawlerEmptyHtmlSelector = "crawler html selector cannot be empty"
	CrawlerFilledPagesBodies = "crawler filled pages bodies needs to be empty"
	CrawlerClosingPageError  = "crawler did not close the page properly"
)

const (
	NoCrawlersInitialized = "no crawlers were initialized"
)

const (
	JSONSerializationFailed = "unable to serialize JSON:"
	FileOpenError           = "unable to open file:"
	FileWriteError          = "unable to write file:"
	HttpFetchError          = "unable to fetch URL:"
)
