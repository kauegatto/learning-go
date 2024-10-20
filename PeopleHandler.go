package main

import (
	"crud_http/stores"
	"encoding/json"
	"fmt"
	"net/http"
)

func handlePeople(store stores.PersonStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			person, found, error := store.GetByName(name)

			if !found {
				fmt.Printf("Did not find person %s\n", name)
				panic("todo")
			}
			if error != nil {
				panic("todo")
			}
			jsonKaue, _ := json.Marshal(person)
			w.Write(jsonKaue)
		},
	)
}
