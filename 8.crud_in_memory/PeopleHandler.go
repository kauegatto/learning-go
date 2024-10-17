package main

import (
	"fmt"
	"net/http"
)

type PersonCrudServer struct{}

func handlePeople() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world")
		},
	)
}
