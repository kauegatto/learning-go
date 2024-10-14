package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppServer(t *testing.T) {
	server := &PersonCrudServer{}

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	got := response.Body.String()

	if got != "Hello, world" {
		t.Errorf("expected %s got %s", "Hello, world", got)
	}
}
