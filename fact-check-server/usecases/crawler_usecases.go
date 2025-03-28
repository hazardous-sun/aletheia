package usecases

import (
	custom_errors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/repositories"
	"encoding/json"
	"fmt"
	"os"
)

type CrawlerUsecase struct {
	crawlerRepository repositories.CrawlerRepository
}

func NewCrawlerUsecase(crawlerRepository repositories.CrawlerRepository) CrawlerUsecase {
	return CrawlerUsecase{
		crawlerRepository: crawlerRepository,
	}
}

func (cu *CrawlerUsecase) Crawl(newsOutlets []models.NewsOutlet, pagesToVisit int, query string) {
	var crawlers []models.Crawler

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
		crawlers = append(crawlers, newCrawler)
	}

	for _, crawler := range crawlers {
		custom_errors.CustomLog(
			fmt.Sprintf("initializing crawler %d", crawler.Id),
			custom_errors.InfoLevel,
		)
	}

	var haltedCrawlers []models.Crawler
	for {
		for i, crawler := range crawlers {
			if crawler.Status != custom_errors.CrawlerRunning {
				haltedCrawlers = append(haltedCrawlers, crawler)
				crawlers = append(crawlers[:i], crawlers[i+1:]...)
			}
		}
		if len(crawlers) == 0 {
			break
		}
	}

	// Serialize crawlers to JSON
	jsonData, err := json.MarshalIndent(crawlers, "", "  ")
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
