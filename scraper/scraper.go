package scraper

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type Scraper struct {
	url       string
	collector *colly.Collector
}

func NewScraper(url string) *Scraper {
	return &Scraper{
		url:       url,
		collector: colly.NewCollector(),
	}
}

var (
	ErrorWebPage    = errors.New("Webpage is down!")
	ErrorInvalidURL = errors.New("Invalid URL!")
)

func (s *Scraper) GetSummary() (w WebPageSummary, err error) {
	webPage, err := s.GetInfo()
	if err != nil {
		return
	}
	w.HTMLVersion = webPage.HTMLVersion
	w.Title = webPage.Title
	w.IsLogIn = webPage.IsLogIn
	w.InternalLinksCount = len(webPage.InternalLinks)
	w.ExternalLinksCount = len(webPage.ExternalLinks)
	w.InAccessibleLinksCount = len(webPage.InAccessibleLinks)
	w.HeadingsCount.Heading1Count = len(webPage.Headings.Heading1)
	w.HeadingsCount.Heading2Count = len(webPage.Headings.Heading2)
	w.HeadingsCount.Heading3Count = len(webPage.Headings.Heading3)
	w.HeadingsCount.Heading4Count = len(webPage.Headings.Heading4)
	w.HeadingsCount.Heading5Count = len(webPage.Headings.Heading5)
	w.HeadingsCount.Heading6Count = len(webPage.Headings.Heading6)
	return
}

func (s *Scraper) GetInfo() (w Webpage, webErr error) {

	parsedURL, err := url.Parse(s.url)
	if err != nil {
		webErr = ErrorInvalidURL
		return
	}

	var hasForm, hasPassword, hasSignInTitle bool
	s.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" {
			return
		}

		if link[0] == '/' || strings.Contains(link, parsedURL.Hostname()) {
			if link[0] == '/' {
				link = s.url + link
			}
			w.InternalLinks = append(w.InternalLinks, link)
			return
		}
		w.ExternalLinks = append(w.ExternalLinks, link)
	})

	s.collector.OnHTML("head title", func(e *colly.HTMLElement) { // Titlee
		w.Title = e.Text
		if !hasSignInTitle && (strings.EqualFold(w.Title, "sign in") || strings.EqualFold(w.Title, "log in")) {
			hasSignInTitle = true
		}
	})

	s.collector.OnHTML("form", func(e *colly.HTMLElement) {
		hasForm = true
	})

	s.collector.OnHTML("h1", func(e *colly.HTMLElement) {
		w.Headings.Heading1 = append(w.Headings.Heading1, e.Text)
	})

	s.collector.OnHTML("h2", func(e *colly.HTMLElement) {
		w.Headings.Heading2 = append(w.Headings.Heading2, e.Text)
	})

	s.collector.OnHTML("h3", func(e *colly.HTMLElement) {
		w.Headings.Heading3 = append(w.Headings.Heading3, e.Text)
	})

	s.collector.OnHTML("h4", func(e *colly.HTMLElement) {
		w.Headings.Heading4 = append(w.Headings.Heading4, e.Text)
	})

	s.collector.OnHTML("h5", func(e *colly.HTMLElement) {
		w.Headings.Heading5 = append(w.Headings.Heading5, e.Text)
	})

	s.collector.OnHTML("h6", func(e *colly.HTMLElement) {
		w.Headings.Heading6 = append(w.Headings.Heading6, e.Text)
	})

	s.collector.OnHTML("input", func(e *colly.HTMLElement) {
		if !hasPassword && e.Attr("type") == "password" {
			hasPassword = true
		}
	})

	s.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	s.collector.OnError(func(_ *colly.Response, err error) {
		webErr = fmt.Errorf("%w : %v", ErrorWebPage, err)
		log.Println("Something went wrong:", err)
	})

	s.collector.OnResponse(func(r *colly.Response) {
		w.HTMLVersion = s.getHTMLVersion(string(r.Body))
	})

	s.collector.Visit(s.url)
	w.IsLogIn = s.isLogIn(s.url) || (hasForm && (hasSignInTitle || hasPassword))
	w.fetchInAccessibleLinks()
	return
}

func (s *Scraper) getHTMLVersion(response string) string {
	if strings.Contains(response, "<!DOCTYPE html>") || strings.Contains(response, "<!doctype html>") {
		return "HTML 5"
	}
	if strings.EqualFold(response, "HTML 4.01") || strings.Contains(response, "html 4.01") {
		return "HTML 4.01"
	}
	if strings.Contains(response, "HTML 2.0") || strings.Contains(response, "html 2.0") {
		return "HTML 2.0"
	}
	if strings.Contains(response, "HTML 3.2") || strings.Contains(response, "html 3.2") {
		return "HTML 3.2"
	}
	if strings.Contains(response, "XHTML Basic 1.0") || strings.Contains(response, "XHTML Basic 1.0") {
		return "XHTML Basic 1.0"
	}
	if strings.Contains(response, "XHTML 1.0") || strings.Contains(response, "xhtml 1.0") {
		return "XHTML 1.0"
	}

	return ""
}

func (s *Scraper) isLogIn(link string) bool {
	if strings.Index(link, "/login") > -1 || strings.Index(link, "/log-in") > -1 || strings.Index(link, "/log_in") > -1 {
		return true
	}
	if strings.Index(link, "/signin") > -1 || strings.Index(link, "/sign-in") > -1 || strings.Index(link, "/sign_in") > -1 {
		return true
	}
	return false
}
