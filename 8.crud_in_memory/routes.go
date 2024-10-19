package main

import "net/http"

func AddRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/people", handlePeople)
	mux.HandleFunc("/people/", handlePeople)
}

/*
Keen viewers may have picked up something odd.
homeHandler() doesn't have ServeHTTP() defined anywhere.
So this must not be a handler, right?

Yesn't.

You are right, homeHandler isn't technically a handler.
However, it is something else; a HandlerFunc.

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
*/
