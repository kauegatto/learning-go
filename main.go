package main

import (
	"crud_http/models"
	"crud_http/stores"
	"log"
	"net/http"
)

func NewServer(
	personStore *stores.PersonStore,
) http.Handler {
	mux := http.NewServeMux()
	AddRoutes(mux, personStore)
	return mux
}

func main() {
	personStore := stores.NewPersonStore()
	personStore.Upsert(models.Person{
		Name:    "kaue",
		BornAt:  "02/06/2003",
		Address: "rua sao vicente",
	})
	personStore.Upsert(models.Person{
		Name:    "douglas",
		BornAt:  "02/07/2003",
		Address: "rua cubatao",
	})

	server := NewServer(personStore)
	log.Fatal(http.ListenAndServe(":8080", server))
}
