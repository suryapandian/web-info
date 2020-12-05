package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"net/http"

	"github.com/suryapandian/web-info/scraper"
)

func setScraperRoutes(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/info", getInfo)
		r.Get("/summary", getSummary)
	})
}

type Req struct {
	URL string `json:"url"`
}

func getSummary(w http.ResponseWriter, r *http.Request) {
	req := Req{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
		return
	}

	if badRequestIfNotMandatoryParams("url", req.URL, w) {
		return
	}

	webPageSummary, err := scraper.NewScraper(req.URL).GetSummary()
	if errors.Is(err, scraper.ErrorWebPage) {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
		return
	}

	switch err {
	case nil:
		writeJSONStruct(webPageSummary, http.StatusOK, w)
	case scraper.ErrorInvalidURL:
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
	default:
		writeJSONMessage(err.Error(), http.StatusInternalServerError, w)

	}

}

func getInfo(w http.ResponseWriter, r *http.Request) {
	req := Req{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
		return
	}

	if badRequestIfNotMandatoryParams("url", req.URL, w) {
		return
	}

	webPage, err := scraper.NewScraper(req.URL).GetInfo()
	if errors.Is(err, scraper.ErrorWebPage) {
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
		return
	}

	switch err {
	case nil:
		writeJSONStruct(webPage, http.StatusOK, w)
	case scraper.ErrorInvalidURL:
		writeJSONMessage(err.Error(), http.StatusBadRequest, w)
	default:
		writeJSONMessage(err.Error(), http.StatusInternalServerError, w)

	}

}
