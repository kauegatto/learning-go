package main

import (
	"crud_http/models"
	"encoding/json"
	"net/http"
)

type PersonCrudServer struct{}

func handlePeople(w http.ResponseWriter, r *http.Request) {
	kaue := models.Person{
		Name:    "kaue",
		BornAt:  "02/06/2003",
		Address: "rua sao vicente",
	}

	jsonKaue, _ := json.Marshal(kaue)
	w.Write(jsonKaue)
}
