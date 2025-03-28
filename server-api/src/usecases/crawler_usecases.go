package usecases

import (
	"encoding/json"
	custom_errors2 "fact-checker-server/src/errors"
	models2 "fact-checker-server/src/models"
	"fact-checker-server/src/repositories"
	"fmt"
	"log"
	"os"
	"time"
)

type CrawlerUsecase struct {
}

func NewCrawlerUsecase() CrawlerUsecase {
	return CrawlerUsecase{}
}

func (cu *CrawlerUsecase) Crawl(newsOutlets []models2.NewsOutlet, pagesToVisit int, query string) {
	var crawlersRepositories []repositories.CrawlerRepository

	// Generate the crawlers for each news outlet returned from the database
	for i, newsOutlet := range newsOutlets {
		newCrawler := models2.Crawler{
			Id:           i + 1,
			PagesToVisit: pagesToVisit,
			Query:        query,
			QueryUrl:     newsOutlet.QueryUrl,
			HtmlSelector: newsOutlet.HtmlSelector,
			Status:       custom_errors2.CrawlerReady,
			PagesBodies:  make([]string, 0),
		}
		crawlersRepositories = append(crawlersRepositories, repositories.NewCrawlerRepository(newCrawler))
	}

	// Initialize the crawlers
	for _, crawlerRepository := range crawlersRepositories {
		custom_errors2.Log(
			fmt.Sprintf("initializing crawler %d", crawlerRepository.Crawler.Id),
			custom_errors2.InfoLevel,
		)
		go crawlerRepository.Crawl()
	}

	// Check for status until all the crawlers finish
	var haltedCrawlers []models2.Crawler
	for {
		for i, crawlerRepository := range crawlersRepositories {
			if crawlerRepository.Crawler.Status != custom_errors2.CrawlerRunning {
				haltedCrawlers = append(haltedCrawlers, crawlerRepository.Crawler)
				crawlersRepositories = append(crawlersRepositories[:i], crawlersRepositories[i+1:]...)
			}
		}
		if len(crawlersRepositories) == 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}

	// Saving the results

	// Serialize the slice to JSON
	jsonData, err := json.MarshalIndent(haltedCrawlers, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing to JSON: %v", err)
	}

	// Open the file in append mode, create it if it doesn't exist, and set write permissions
	file, err := os.OpenFile("results", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Add a newline after appending data for better readability
	if _, err := file.WriteString("\n"); err != nil {
		log.Fatalf("Error writing newline to file: %v", err)
	}
}
