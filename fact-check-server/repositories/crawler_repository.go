package repositories

import (
	custom_errors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type CrawlerRepository struct {
	crawler models.Crawler
}

func (cr *CrawlerRepository) Crawl() error {
	// crawler.QueryUrl should not be empty
	if cr.crawler.QueryUrl == "" {
		custom_errors.CustomLog(
			"crawler failed because it was initialized without an URL to query",
			custom_errors.WarningLevel,
		)
		return errors.New(custom_errors.CrawlerEmptyQueryUrl)
	}

	// crawler.Query should not be empty
	if cr.crawler.Query == "" {
		custom_errors.CustomLog(
			fmt.Sprintf("crawler querying %s failed because query was empty", cr.crawler.QueryUrl),
			custom_errors.WarningLevel,
		)
		return errors.New(custom_errors.CrawlerEmptyQuery)
	}

	// crawler.HtmlSelector should not be empty
	if cr.crawler.HtmlSelector == "" {
		// crawler.QueryUrl should not be empty
		if cr.crawler.QueryUrl == "" {
			custom_errors.CustomLog(
				fmt.Sprintf("crawler querying %s failed because HTML selector was empty", cr.crawler.QueryUrl),
				custom_errors.WarningLevel,
			)
			return errors.New(custom_errors.CrawlerEmptyQueryUrl)
		}
		return errors.New(custom_errors.CrawlerEmptyHtmlSelector)
	}

	// crawler.PagesBodies should be empty
	if len(cr.crawler.PagesBodies) > 0 {
		// crawler.QueryUrl should not be empty
		if cr.crawler.QueryUrl == "" {
			custom_errors.CustomLog(
				fmt.Sprintf("crawler querying %s failed because page bodies was initialized with values already maintained", cr.crawler.QueryUrl),
				custom_errors.WarningLevel,
			)
			return errors.New(custom_errors.CrawlerEmptyQueryUrl)
		}
		return errors.New(custom_errors.CrawlerFilledPagesBodies)
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

		// Extract the main article link
		link := e.Attr("href")
		if link != "" {
			results = append(results, link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		custom_errors.CustomLog(fmt.Sprintf("Visiting: %s", r.URL.String()), custom_errors.InfoLevel)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	if err := c.Visit(searchURL); err != nil {
		fmt.Println("Visit error:", err)
	}

	// Fetch and save the body content of each link
	for i, link := range results {
		if !strings.HasPrefix(link, "http") {
			link = "https:" + link // Ensure the link has a valid scheme
		}

		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", link, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading body from %s: %v\n", link, err)
			continue
		}

		// Write the body to a file
		filename := fmt.Sprintf("result%d", i+1)
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Saved body from %s to %s\n", link, filename)
	}
}
