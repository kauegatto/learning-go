package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PersonCrudServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
