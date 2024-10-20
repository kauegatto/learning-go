package main

import (
	"crud_http/stores"
	"net/http"
)

func AddRoutes(
	mux *http.ServeMux,
	personStore *stores.PersonStore,
) {
	mux.Handle("/people/{name}", handlePeople(*personStore))
}
