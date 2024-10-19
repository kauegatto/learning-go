package main

import (
	"crud_http/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppServer(t *testing.T) {
	server := NewServer()
	request := httptest.NewRequest(http.MethodGet, "/people/kaue", nil)
	response := httptest.NewRecorder()
	want := models.Person{
		Name:    "kaue",
		BornAt:  "02/06/2003",
		Address: "rua sao vicente",
	}

	server.ServeHTTP(response, request)

	var got models.Person

	err := json.Unmarshal([]byte(response.Body.String()), &got)

	if err != nil {
		panic(err)
	}

	if got != want {
		t.Errorf("expected %s got %s", want, got)
	}
}
