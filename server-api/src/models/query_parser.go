package models

import (
	"aletheia-server/src/errors"
	"fmt"
	"net/url"
	"strings"
)

type QueryParser struct {
	NewsOutletName string
	QueryParam     string
	QueryUrl       string
}

func (qp *QueryParser) Parse() string {
	qp.NewsOutletName = strings.TrimSpace(qp.NewsOutletName)
	if qp.NewsOutletName == "" {
		server_errors.Log(server_errors.EmptyNewsOutletName, server_errors.ErrorLevel)
		return ""
	}

	server_errors.Log(fmt.Sprintf("Generating query to '%s' crawler", qp.NewsOutletName), server_errors.InfoLevel)

	qp.QueryParam = strings.TrimSpace(qp.QueryParam)
	if qp.QueryParam == "" {
		server_errors.Log(server_errors.EmptyQueryParam, server_errors.ErrorLevel)
		return ""
	}

	// Replace spaces with '+' and URL encode special characters
	encodedQuery := strings.ReplaceAll(url.QueryEscape(qp.QueryParam), "%20", "+")

	qp.QueryUrl = strings.TrimSpace(qp.QueryUrl)
	if qp.QueryUrl == "" {
		server_errors.Log(server_errors.EmptyQueryUrl, server_errors.ErrorLevel)
		return ""
	}

	finalQuery := strings.ReplaceAll(qp.QueryUrl, "QUERY_HERE", encodedQuery)
	server_errors.Log(fmt.Sprintf("QueryParam to %s generated: %s", qp.NewsOutletName, finalQuery), server_errors.InfoLevel)
	return finalQuery
}
