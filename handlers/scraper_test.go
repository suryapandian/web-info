package handlers

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSummary(t *testing.T) {
	var testCases = []struct {
		desc               string
		payload            []byte
		expectedStatusCode int
	}{
		{
			"no payload",
			nil,
			http.StatusBadRequest,
		},
		{
			"invalid json",
			[]byte(`{"key":}`),
			http.StatusBadRequest,
		},
		{
			"valid payload",
			[]byte(fmt.Sprintf(`{
				"url": "%s"
			}`, "https://www.google.com")),
			http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := assert.New(t)
			r := httptest.NewRequest(http.MethodGet, "/summary", bytes.NewReader(testCase.payload))
			w := httptest.NewRecorder()
			GetRouter().ServeHTTP(w, r)
			response := w.Result()
			a.Equal(testCase.expectedStatusCode, response.StatusCode)
		})
	}

}
