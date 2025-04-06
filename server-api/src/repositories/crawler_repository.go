package repositories

import (
	"aletheia-server/src/errors"
	"aletheia-server/src/models"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
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

/*
1. Analyze QueryUrl
	1. Replace KEYWORDS_HERE1 by query value
2. Analyze HtmlSelector
	1. Replace N_FIRST_... with the number of the first iteration
		- Example: N_FIRST_2 should be converted to 2, N_FIRST_13 should be converted to 13...
3. Crawl QueryUrl
	1. While len(resultsFound) < PagesToVisit, iterate over the resulting page searching for links
*/

func (cr *CrawlerRepository) Crawl() {
	cr.Crawler.Status = server_errors.CrawlerRunning

	// crawler.QueryUrl should not be empty
	if cr.Crawler.QueryUrl == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because it was initialized without an URL to query", cr.Crawler.Id),
			server_errors.WarningLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerEmptyQueryUrl
		return
	}

	// crawler.Query should not be empty
	if cr.Crawler.Query == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because query was empty", cr.Crawler.Id),
			server_errors.WarningLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerEmptyQuery
		return
	}

	// crawler.HtmlSelector should not be empty
	if cr.Crawler.HtmlSelector == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because HTML selector was empty", cr.Crawler.Id),
			server_errors.WarningLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerEmptyHtmlSelector
		return
	}

	// crawler.PagesBodies should be empty
	if len(cr.Crawler.PagesBodies) > 0 {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because its page bodies was initialized with values already maintained", cr.Crawler.Id),
			server_errors.WarningLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerFilledPagesBodies
		return
	}

	searchURL := strings.ReplaceAll(cr.Crawler.QueryUrl, "KEYWORDS_HERE", cr.Crawler.Query)
	searchURL = url.QueryEscape(searchURL)
	searchURL = strings.ReplaceAll(searchURL, "+", "%20")
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
			fmt.Sprintf("crawler %d visiting: %s", cr.Crawler.Id, searchURL),
			server_errors.InfoLevel)
	})

	c.OnError(func(_ *colly.Response, err error) {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed: %s", cr.Crawler.Id, searchURL),
			server_errors.ErrorLevel)
	})

	if err := c.Visit(searchURL); err != nil {
		server_errors.Log(
			fmt.Sprintf("crawler %d visit error: %s", cr.Crawler.Id, searchURL),
			server_errors.ErrorLevel)
	}

	// Fetch and save the body content of each link
	for _, link := range results {
		collectCandidateBody(cr, link)
	}
	cr.Crawler.Status = server_errors.CrawlerSucceeded
}

func collectCandidateBody(cr *CrawlerRepository, link string) {
	if !strings.HasPrefix(link, "http") {
		link = "https:" + link // Ensure the link has a valid scheme
	}

	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", link, err)
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

	// TODO the bodies are not being correctly stored here, the sites are not being visited
	// Store the body
	candidateBody := string(body)
	cr.Crawler.PagesBodies = append(cr.Crawler.PagesBodies, candidateBody)

	server_errors.Log(
		fmt.Sprintf("added %s to crawler %d pagebodies", link, cr.Crawler.Id),
		server_errors.InfoLevel,
	)
}

// Replaces `N_FIRST_(\d+)` with the current iteration
func configHtmlSelector(input string) string {
	re := regexp.MustCompile(`N_FIRST_(\d+)`)
	return re.ReplaceAllString(input, "$1")
}
