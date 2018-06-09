package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New creates and starts a new server.
func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/{albumName}/{fileName}", GetImage).Methods("GET")
	router.HandleFunc("/{albumName}/{fileName}", CreateImage).Methods("POST")

	return router
}
