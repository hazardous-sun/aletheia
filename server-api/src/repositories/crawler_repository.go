package repositories

import (
	"aletheia-server/src/errors"
	"aletheia-server/src/models"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type CrawlerRepository struct {
	Crawler models.Crawler
}

func NewCrawlerRepository(crawler models.Crawler) CrawlerRepository {
	return CrawlerRepository{
		Crawler: crawler,
	}
}

func (cr *CrawlerRepository) Crawl() {
	cr.Crawler.Status = server_errors.CrawlerRunning

	if badCrawler(&cr.Crawler) {
		return
	}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var results []string

	// Look for URLs that should be visited
	c.OnHTML(cr.Crawler.HtmlSelector, func(e *colly.HTMLElement) {
		if len(results) >= cr.Crawler.PagesToVisit {
			return
		}

		// Extract the URL
		link := e.Attr("href")
		if link != "" {
			results = append(results, link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		server_errors.Log(
			fmt.Sprintf("crawler %d visiting: %s", cr.Crawler.Id, cr.Crawler.Query),
			server_errors.InfoLevel)
	})

	c.OnError(func(_ *colly.Response, err error) {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed: %s", cr.Crawler.Id, cr.Crawler.Query),
			server_errors.ErrorLevel)
	})

	if err := c.Visit(cr.Crawler.Query); err != nil {
		server_errors.Log(
			fmt.Sprintf("crawler %d visit error: %s", cr.Crawler.Id, cr.Crawler.Query),
			server_errors.ErrorLevel)
		cr.Crawler.Status = err.Error()
		return
	}

	// Fetch and save the body content of each link
	for _, link := range results {
		collectCandidateBody(cr, link)
	}
	cr.Crawler.Status = server_errors.CrawlerSucceeded
}

func badCrawler(crawler *models.Crawler) bool {
	if crawler.Query == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because it was initialized without a query", crawler.Id),
			server_errors.ErrorLevel,
		)
		crawler.Status = server_errors.CrawlerEmptyQueryUrl
		return true
	}

	// crawler.HtmlSelector should not be empty
	if crawler.HtmlSelector == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because HTML selector was empty", crawler.Id),
			server_errors.ErrorLevel,
		)
		crawler.Status = server_errors.CrawlerEmptyHtmlSelector
		return true
	}

	// crawler.PagesBodies should be empty
	if len(crawler.PagesBodies) > 0 {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because its page bodies was initialized with values already maintained", crawler.Id),
			server_errors.ErrorLevel,
		)
		crawler.Status = server_errors.CrawlerFilledPagesBodies
		return true
	}

	return false
}

func collectCandidateBody(cr *CrawlerRepository, link string) {
	if !strings.HasPrefix(link, "http") {
		link = "https:" + link // Ensure the link has a valid scheme
	}

	resp, err := http.Get(link)
	if err != nil {
		server_errors.Log(fmt.Sprintf("%s %s ->", server_errors.HttpFetchError, link), server_errors.ErrorLevel)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			server_errors.Log(
				server_errors.CrawlerClosingPageError,
				server_errors.WarningLevel,
			)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		server_errors.Log(
			fmt.Sprintf("unable to read body from %s: %v", link, err),
			server_errors.ErrorLevel,
		)
		return
	}

	// Store the body
	candidateBody := string(body)
	cr.Crawler.PagesBodies = append(cr.Crawler.PagesBodies, candidateBody)

	server_errors.Log(
		fmt.Sprintf("added %s to crawler %d pagebodies", link, cr.Crawler.Id),
		server_errors.InfoLevel,
	)
}
