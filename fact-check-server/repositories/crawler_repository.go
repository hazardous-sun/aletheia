package repositories

import (
	custom_errors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type CrawlerRepository struct {
	crawler models.Crawler
}

func (cr *CrawlerRepository) Crawl() {
	// crawler.QueryUrl should not be empty
	if cr.crawler.QueryUrl == "" {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d failed because it was initialized without an URL to query", cr.crawler.Id),
			custom_errors.WarningLevel,
		)
		cr.crawler.Status = custom_errors.CrawlerEmptyQueryUrl
		return
	}

	// crawler.Query should not be empty
	if cr.crawler.Query == "" {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d failed because query was empty", cr.crawler.Id),
			custom_errors.WarningLevel,
		)
		cr.crawler.Status = custom_errors.CrawlerEmptyQuery
		return
	}

	// crawler.HtmlSelector should not be empty
	if cr.crawler.HtmlSelector == "" {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d failed because HTML selector was empty", cr.crawler.Id),
			custom_errors.WarningLevel,
		)
		cr.crawler.Status = custom_errors.CrawlerEmptyHtmlSelector
		return
	}

	// crawler.PagesBodies should be empty
	if len(cr.crawler.PagesBodies) > 0 {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d failed because its page bodies was initialized with values already maintained", cr.crawler.Id),
			custom_errors.WarningLevel,
		)
		cr.crawler.Status = custom_errors.CrawlerFilledPagesBodies
		return
	}

	searchURL := strings.ReplaceAll(cr.crawler.QueryUrl, "KEYWORDS_HERE", cr.crawler.Query)
	searchURL = url.QueryEscape(searchURL)
	searchURL = strings.ReplaceAll(searchURL, "+", "%20")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var results []string

	// Look for URLs that should be visited
	c.OnHTML(cr.crawler.HtmlSelector, func(e *colly.HTMLElement) {
		if len(results) >= cr.crawler.PagesToVisit {
			return
		}

		// Extract the URL
		link := e.Attr("href")
		if link != "" {
			results = append(results, link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d visiting: %s", cr.crawler.Id, searchURL),
			custom_errors.InfoLevel)
	})

	c.OnError(func(_ *colly.Response, err error) {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d failed: %s", cr.crawler.Id, searchURL),
			custom_errors.ErrorLevel)
	})

	if err := c.Visit(searchURL); err != nil {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler %d visit error: %s", cr.crawler.Id, searchURL),
			custom_errors.ErrorLevel)
	}

	// Fetch and save the body content of each link
	for _, link := range results {
		if !strings.HasPrefix(link, "http") {
			link = "https:" + link // Ensure the link has a valid scheme
		}

		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", link, err)
			continue
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				custom_errors.CustomLog(
					custom_errors.CrawlerClosingPageError,
					custom_errors.WarningLevel,
				)
			}
		}(resp.Body)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			custom_errors.CustomLog(
				fmt.Sprintf("unable to read body from %s: %v", link, err),
				custom_errors.ErrorLevel,
			)
			continue
		}

		// Store the body
		cr.crawler.PagesBodies = append(cr.crawler.PagesBodies, string(body))

		custom_errors.CustomLog(
			fmt.Sprintf("added %s crawler %d pagebodies", link),
			custom_errors.ErrorLevel,
		)
	}
	cr.crawler.Status = custom_errors.CrawlerSucceeded
}
