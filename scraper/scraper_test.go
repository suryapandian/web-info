package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {

	html := `
		<!DOCTYPE html>
		<html>
		<head>
		<title>Test Page</title>
		</head>
		<body>
		<h1>Hello World</h1>
		<a href="https://www.google.com/" />
		<a href="invalid/" />
		</body>
		</html>
	`
	//Setup a mock  server from which we can simulate
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(html))
	}))
	defer mockServer.Close()
	scraper := NewScraper(mockServer.URL)

	a := assert.New(t)
	summary, err := scraper.GetSummary()
	a.NoError(err)
	a.Equal("Test Page", summary.Title)
	a.Equal(1, summary.HeadingsCount.Heading1Count)
	a.Equal(0, summary.InternalLinksCount)
	a.Equal(2, summary.ExternalLinksCount)
	a.Equal(1, summary.InAccessibleLinksCount)
	a.Equal(false, summary.IsLogIn)
}
