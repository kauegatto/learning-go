package main

import (
	"fmt"
	"net/http"
)

type PersonCrudServer struct{}

func (server *PersonCrudServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}
