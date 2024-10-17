package main

import (
	"log"
	"net/http"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	AddRoutes(mux)
	return mux
}

func main() {
	// server := &PersonCrudServer{}
	server := NewServer()
	log.Fatal(http.ListenAndServe(":8080", server))
}

/*
NewServer is a big constructor that takes in all dependencies as arguments
It returns an http.Handler if possible, which can be a dedicated type for more complex situations
It usually configures its own muxer and calls out to routes.go


*/
