package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppServer(t *testing.T) {
	server := NewServer()
	request := httptest.NewRequest(http.MethodGet, "/people/kaue", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	got := response.Body.String()
	want := "Hello, world"

	if got != want {
		t.Errorf("expected %s (len: %d) got %s (len: %d)",
			want, len(want),
			got, len(got))
	}
}
