package main

import "net/http"

func AddRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/people/{name}", handlePeople)
}
