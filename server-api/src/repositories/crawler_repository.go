package repositories

import (
	"aletheia-server/src/errors"
	"aletheia-server/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	if cr.badCrawler() {
		return
	}

	// Get the initial page content
	resp, err := http.Get(cr.Crawler.Query)
	if err != nil {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed to fetch initial page: %v", cr.Crawler.Id, err),
			server_errors.ErrorLevel)
		cr.Crawler.Status = err.Error()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed to read initial page: %v", cr.Crawler.Id, err),
			server_errors.ErrorLevel)
		cr.Crawler.Status = err.Error()
		return
	}

	// Send HTML content to AI analyzer to get links
	links, err := getLinksFromAI(string(body))
	if err != nil {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed to get links from AI: %v", cr.Crawler.Id, err),
			server_errors.ErrorLevel)
		cr.Crawler.Status = err.Error()
		return
	}

	// Limit the number of pages to visit
	if len(links) > cr.Crawler.PagesToVisit {
		links = links[:cr.Crawler.PagesToVisit]
	}

	// Fetch and save the body content of each link
	for _, link := range links {
		cr.collectCandidateBody(link)
	}
	cr.Crawler.Status = server_errors.CrawlerSucceeded
}

func (cr *CrawlerRepository) badCrawler() bool {
	if cr.Crawler.Query == "" {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because it was initialized without a query", cr.Crawler.Id),
			server_errors.ErrorLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerEmptyQueryUrl
		return true
	}

	// crawler.PagesBodies should be empty
	if len(cr.Crawler.PagesBodies) > 0 {
		server_errors.Log(
			fmt.Sprintf("crawler %d failed because its page bodies was initialized with values already maintained", cr.Crawler.Id),
			server_errors.ErrorLevel,
		)
		cr.Crawler.Status = server_errors.CrawlerFilledPagesBodies
		return true
	}

	return false
}

func (cr *CrawlerRepository) collectCandidateBody(link string) {
	if !strings.HasPrefix(link, "http") {
		link = "https://" + link // Ensure the link has a valid scheme
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

func getLinksFromAI(htmlContent string) ([]string, error) {
	htmlContent = strings.ReplaceAll(htmlContent, "\"", "'") // Escape quotes
	htmlContent = strings.ReplaceAll(htmlContent, "\n", "")  // Remove newlines

	// Prepare request to AI analyzer
	requestBody, err := json.Marshal(map[string]string{
		"html_content": htmlContent,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	aiServiceURL := os.Getenv("AI_ANALYZER_URL")
	if aiServiceURL == "" {
		aiServiceURL = "http://localhost:7654" // Default fallback
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/getLinks", aiServiceURL),
		"application/json",
		bytes.NewBuffer(requestBody),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to request links from AI: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AI analyzer returned status %d", resp.StatusCode)
	}

	// Define the correct structure to match the response
	var links []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&links); err != nil {
		server_errors.Log(fmt.Sprintf("resp.Body = %s", resp.Body), server_errors.WarningLevel)
		return nil, fmt.Errorf("failed to decode AI response: %v", err)
	}

	// Extract URLs from the links
	var urls []string
	for _, link := range links {
		urls = append(urls, link.URL)
	}

	return urls, nil
}
