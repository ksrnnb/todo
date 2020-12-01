package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccessTopPage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", top)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v, expected 200", writer.Code)
	}

	isTop := strings.Contains(writer.Body.String(), "top")

	if !isTop {
		t.Error(`Cannot see "top page" at path "/"`)
	}
}
