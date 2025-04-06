package server_errors

import (
	server_errors "aletheia-server/src/errors"
	"testing"
)

func TestCrawlerStatusConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		want     string
	}{
		{
			name:     "CrawlerReady",
			constant: server_errors.CrawlerReady,
			want:     "crawler is ready",
		},
		{
			name:     "CrawlerRunning",
			constant: server_errors.CrawlerRunning,
			want:     "crawler is running",
		},
		{
			name:     "CrawlerSucceeded",
			constant: server_errors.CrawlerSucceeded,
			want:     "crawler successfully crawled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.want {
				t.Errorf("got %q, want %q", tt.constant, tt.want)
			}
		})
	}
}

func TestCrawlerErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		want     string
	}{
		{
			name:     "CrawlerEmptyQuery",
			constant: server_errors.CrawlerEmptyQuery,
			want:     "crawler query cannot be empty",
		},
		{
			name:     "CrawlerEmptyQueryUrl",
			constant: server_errors.CrawlerEmptyQueryUrl,
			want:     "crawler query url cannot be empty",
		},
		{
			name:     "CrawlerEmptyHtmlSelector",
			constant: server_errors.CrawlerEmptyHtmlSelector,
			want:     "crawler html selector cannot be empty",
		},
		{
			name:     "CrawlerFilledPagesBodies",
			constant: server_errors.CrawlerFilledPagesBodies,
			want:     "crawler filled pages bodies needs to be empty",
		},
		{
			name:     "CrawlerClosingPageError",
			constant: server_errors.CrawlerClosingPageError,
			want:     "crawler did not close the page properly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.want {
				t.Errorf("got %q, want %q", tt.constant, tt.want)
			}
		})
	}
}
