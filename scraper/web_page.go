package scraper

import (
	//"net/http"
	"net/url"
	"sync"
)

type Webpage struct {
	HTMLVersion       string   `json:"htmlVersion"`
	Title             string   `json:"title"`
	Headings          Headings `json:"headings"`
	InternalLinks     []string `json:"internalLinks"`
	ExternalLinks     []string `json:"externalLinks"`
	InAccessibleLinks []string `json:"inAccessibleLinks"`
	IsLogIn           bool     `json:"isLogIn"`
}

type WebPageSummary struct {
	HTMLVersion            string        `json:"htmlVersion"`
	Title                  string        `json:"title"`
	HeadingsCount          HeadingsCount `json:"headingsCount"`
	InternalLinksCount     int           `json:"internalLinksCount"`
	ExternalLinksCount     int           `json:"externalLinksCount"`
	InAccessibleLinksCount int           `json:"inAccessibleLinksCount"`
	IsLogIn                bool          `json:"isLogIn"`
}

type HeadingsCount struct {
	Heading1Count int `json:"heading1Count"`
	Heading2Count int `json:"heading2Count"`
	Heading3Count int `json:"heading3Count"`
	Heading4Count int `json:"heading4Count"`
	Heading5Count int `json:"heading5Count"`
	Heading6Count int `json:"heading6Count"`
}

type Headings struct {
	Heading1 []string `json:"heading1"`
	Heading2 []string `json:"heading2"`
	Heading3 []string `json:"heading3"`
	Heading4 []string `json:"heading4"`
	Heading5 []string `json:"heading5"`
	Heading6 []string `json:"heading6"`
}

func (w *Webpage) fetchInAccessibleLinks() {
	links := append(w.InternalLinks, w.ExternalLinks...)
	inAccessibleLinks := make(chan string, len(links))

	var wg sync.WaitGroup
	wg.Add(len(links))

	for _, link := range links {
		go func(link string) {
			defer wg.Done()

			_, err := url.ParseRequestURI(link)
			if err != nil {
				inAccessibleLinks <- link
				return
			}

			// Make actual request to validate
			// if _, err := http.Get(link); err != nil {
			// 	inAccessibleLinks <- link
			//}

		}(link)
	}
	wg.Wait()
	close(inAccessibleLinks)

	for link := range inAccessibleLinks {
		w.InAccessibleLinks = append(w.InAccessibleLinks, link)
	}

	return

}
