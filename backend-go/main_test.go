package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	expected := `{"message":"hello world"}`
	if w.Body.String() != expected {
		t.Fatalf("Expected body %s, got %s", expected, w.Body.String())
	}
}