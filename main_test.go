package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
	"bytes"
)

func init() {
	initDb()
}

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

func TestHandInsertBookmarkWithStatusOK(t *testing.T) {
	bookmark := Bookmark{"wercker", "http://wercker.com"}
	
	b, err := json.Marshal(bookmark)
	if err != nil {
		t.Fatalf("Unable to marshal Bookmark")
	}

	request, _ := http.NewRequest("POST", "/new", bytes.NewReader(b))
	response := httptest.NewRecorder()

	insertBookmark(response, request)
	
	body := response.Body.String()	
	if !strings.Contains(body, "{'bookmark':'saved'}") {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "San Francisco", body)
	}
}

func TestHandleIndexReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	
	handleIndex(response, request)
	
	ct := response.HeaderMap["Content-Type"][0]
	if !strings.EqualFold(ct, "application/json") {
		t.Fatalf("Content-Type does not equal 'application/json'")
	}
}

