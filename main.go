package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Cookies handler
	cookies := r.PathPrefix("/cookies").Subrouter()
	cookies.HandleFunc("", cookiesHandler.Store).Methods("POST")
	cookies.HandleFunc("/{id}", cookiesHandler.Load).Methods("GET")

	// Profile handler
	profile := r.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("/{username}/{fields}", profileHandler.Data).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
