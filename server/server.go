package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

// New creates and starts a new server.
func New(addr string) error {
	router := mux.NewRouter()

	router.HandleFunc("/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := vars["fileName"]
		path := path.Join("test-images", fileName)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		file, err := os.Open(path)

		if err != nil {
			fmt.Fprintf(w, "Bad file")
		}

		io.Copy(w, file)
	}).Methods("GET")

	router.HandleFunc("/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bad file")
	}).Methods("POST")

	http.Handle("/", router)

	go http.ListenAndServe(addr, nil)

	return nil
}
