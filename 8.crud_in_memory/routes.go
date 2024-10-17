package main

import "net/http"

func AddRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("/people", handlePeople())
	mux.Handle("/people/", handlePeople())
}
