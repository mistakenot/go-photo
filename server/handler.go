package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

// GetImage returns an image if it exists.
func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	albumName := vars["albumName"]
	path := path.Join("test-images", albumName, fileName)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(w, "Bad file")
	}

	io.Copy(w, file)
}

// CreateImage will create a new image if it doesn't already exist.
func CreateImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	albumName := vars["albumName"]
	path := path.Join("test-images", albumName, fileName)
	_, err := os.Create(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
