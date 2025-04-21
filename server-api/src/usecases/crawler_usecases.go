package usecases

import (
	"aletheia-server/src/errors"
	"aletheia-server/src/models"
	"aletheia-server/src/repositories"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type CrawlerUsecase struct {
}

func NewCrawlerUsecase() CrawlerUsecase {
	return CrawlerUsecase{}
}

func (cu *CrawlerUsecase) Crawl(newsOutlets []models.NewsOutlet, pagesToVisit int, query string) {
	var crawlersRepositories []repositories.CrawlerRepository

	// Generate the crawlers for each news outlet returned from the database
	for i, newsOutlet := range newsOutlets {
		queryParser := models.QueryParser{
			NewsOutletName: newsOutlet.Name,
			QueryParam:     query,
			QueryUrl:       newsOutlet.QueryUrl,
		}
		finalQuery := queryParser.Parse()

		server_errors.Log(fmt.Sprintf("Parsed query '%s' into '%s'", query, finalQuery), server_errors.InfoLevel)

		if finalQuery == "" {
			continue
		}

		newCrawler := models.Crawler{
			Id:           i + 1,
			PagesToVisit: pagesToVisit,
			Query:        finalQuery,
			HtmlSelector: newsOutlet.HtmlSelector,
			Status:       server_errors.CrawlerReady,
			PagesBodies:  make([]string, 0),
		}
		crawlersRepositories = append(crawlersRepositories, repositories.NewCrawlerRepository(newCrawler))
	}

	// Check if at least one crawler was generated
	if len(crawlersRepositories) == 0 {
		server_errors.Log(server_errors.NoCrawlersInitialized, server_errors.ErrorLevel)
		return
	}

	// Initialize Crawlers concurrently
	var wg sync.WaitGroup
	for _, crawlerRepository := range crawlersRepositories {
		wg.Add(1)
		go func(cr repositories.CrawlerRepository) {
			defer wg.Done()
			server_errors.Log(
				fmt.Sprintf("Initializing crawler %d", cr.Crawler.Id),
				server_errors.InfoLevel,
			)
			cr.Crawl()
		}(crawlerRepository)
	}

	// Wait for all crawlers to finish
	wg.Wait()

	// Collect results after all crawlers are done
	haltedCrawlers := make([]models.Crawler, len(crawlersRepositories))
	for i, cr := range crawlersRepositories {
		haltedCrawlers[i] = cr.Crawler
	}

	// Saving the results
	saveResults(haltedCrawlers)
}

func saveResults(crawlers []models.Crawler) {
	// Serialize the slice to JSON
	jsonData, err := json.MarshalIndent(crawlers, "", "  ")
	if err != nil {
		server_errors.Log(fmt.Sprintf("%s %s", server_errors.JSONSerializationFailed, err.Error()), server_errors.ErrorLevel)
		return
	}

	// Open the file in append mode, create it if it doesn't exist, and set write permissions
	file, err := os.OpenFile("results", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		server_errors.Log(fmt.Sprintf("%s %s", server_errors.FileOpenError, err.Error()), server_errors.ErrorLevel)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	if _, err := file.Write(jsonData); err != nil {
		server_errors.Log(fmt.Sprintf("%s %s", server_errors.FileWriteError, err.Error()), server_errors.ErrorLevel)
		return
	}

	// Add a newline after appending data for better readability
	if _, err := file.WriteString("\n"); err != nil {
		server_errors.Log(fmt.Sprintf("%s %s", server_errors.FileWriteError, err.Error()), server_errors.ErrorLevel)
	}
}
