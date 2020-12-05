package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func badRequestIfNotMandatoryParams(key string, value string, w http.ResponseWriter) bool {
	if strings.TrimSpace(value) == "" {
		writeJSONMessage(fmt.Sprintf("%s is mandatory", key), http.StatusBadRequest, w)
		return true
	}
	return false
}
