package usecases

import (
	custom_errors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/repositories"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type CrawlerUsecase struct {
}

func NewCrawlerUsecase(crawlerRepository repositories.CrawlerRepository) CrawlerUsecase {
	return CrawlerUsecase{}
}

func (cu *CrawlerUsecase) Crawl(newsOutlets []models.NewsOutlet, pagesToVisit int, query string) {
	var crawlersRepositories []repositories.CrawlerRepository

	// Generate the crawlers for each news outlet returned from the database
	for i, newsOutlet := range newsOutlets {
		newCrawler := models.Crawler{
			Id:           i + 1,
			PagesToVisit: pagesToVisit,
			Query:        query,
			QueryUrl:     newsOutlet.QueryUrl,
			HtmlSelector: newsOutlet.HtmlSelector,
			Status:       custom_errors.CrawlerReady,
			PagesBodies:  make([]string, 0),
		}
		crawlersRepositories = append(crawlersRepositories, repositories.NewCrawlerRepository(newCrawler))
	}

	// Initialize the crawlers
	for _, crawlerRepository := range crawlersRepositories {
		custom_errors.CustomLog(
			fmt.Sprintf("initializing crawler %d", crawlerRepository.Crawler.Id),
			custom_errors.InfoLevel,
		)
		crawlerRepository.Crawl()
	}

	// Check for status until all the crawlers finish
	var haltedCrawlers []models.Crawler
	for {
		for i, crawlerRepository := range crawlersRepositories {
			if crawlerRepository.Crawler.Status != custom_errors.CrawlerRunning {
				haltedCrawlers = append(haltedCrawlers, crawlerRepository.Crawler)
				crawlersRepositories = append(crawlersRepositories[:i], crawlersRepositories[i+1:]...)
			}
		}
		if len(crawlersRepositories) == 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}

	// Serialize crawlersRepositories to JSON
	jsonData, err := json.MarshalIndent(crawlersRepositories, "", "  ")
	if err != nil {
		custom_errors.CustomLog(
			fmt.Sprintf("unable to serialize to JSON: %v", err),
			custom_errors.ErrorLevel,
		)
	}

	// Write the JSON data to a file named "results"
	err = os.WriteFile("results", jsonData, 0644)
	if err != nil {
		custom_errors.CustomLog(
			fmt.Sprintf("unable to write to file: %v", err),
			custom_errors.ErrorLevel,
		)
	}
}
