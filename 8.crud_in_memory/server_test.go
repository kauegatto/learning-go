package main

import (
	"crud_http/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertEquals(t *testing.T, got any, want any) {
	if got != want {
		t.Errorf("expected %s got %s", want, got)
	}
}
func TestPersonHandler(t *testing.T) {
	server := NewServer()
	t.Run("Get kaue is correct", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/people/kaue", nil)
		response := httptest.NewRecorder()
		want := models.Person{
			Name:    "kaue",
			BornAt:  "02/06/2003",
			Address: "rua sao vicente",
		}

		server.ServeHTTP(response, request)

		var got models.Person

		json.Unmarshal(response.Body.Bytes(), &got)

		assertEquals(t, got, want)
	})
	t.Run("Get kaue is correct", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/people/douglas", nil)
		response := httptest.NewRecorder()
		want := models.Person{
			Name:    "Douglas Reis",
			BornAt:  "02/07/2003",
			Address: "rua cubatao",
		}

		server.ServeHTTP(response, request)

		var got models.Person

		json.Unmarshal(response.Body.Bytes(), &got)

		assertEquals(t, got, want)
	})
}

func TestAppServer(t *testing.T) {

}
