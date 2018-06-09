package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New creates and starts a new server.
func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/{fileName}", GetImage).Methods("GET")
	router.HandleFunc("/{fileName}", CreateImage).Methods("POST")

	return router
}
