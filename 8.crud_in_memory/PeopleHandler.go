package main

import (
	"crud_http/models"
	"encoding/json"
	"net/http"
)

type PersonCrudServer struct{}

func handlePeople(w http.ResponseWriter, r *http.Request) {
	person := r.PathValue("name")
	var found models.Person
	if person == "kaue" {
		found = models.Person{
			Name:    "kaue",
			BornAt:  "02/06/2003",
			Address: "rua sao vicente",
		}
	} else {
		found = models.Person{
			Name:    "Douglas Reis",
			BornAt:  "02/07/2003",
			Address: "rua cubatao",
		}
	}

	jsonKaue, _ := json.Marshal(found)
	w.Write(jsonKaue)
}
