package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New creates and starts a new server.
func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api", GetAlbumOverview).Methods("GET")
	router.HandleFunc("/api/{albumName}", GetAlbum).Methods("GET")
	router.HandleFunc("/api/{albumName}/{fileName}", GetImage).Methods("GET")
	router.HandleFunc("/api/{albumName}/{fileName}", CreateImage).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./test-images")))

	return router
}
