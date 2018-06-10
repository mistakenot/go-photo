package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// GetAlbumOverview returns a summary of all albums.
func GetAlbumOverview(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./test-images")

	if err != nil {

	}

	albums := make([]Album, len(files))
	size := 0

	for i := 0; i < len(files); i++ {
		photos, err := ioutil.ReadDir("./test-images/" + files[i].Name())

		if err != nil {

		}

		albums[i] = Album{
			Name:  files[i].Name(),
			Size:  int(files[i].Size()),
			Count: len(photos)}
		size += int(files[i].Size())
	}

	model := AlbumOverview{
		Albums: albums}

	bytes, err := json.Marshal(model)

	if err != nil {

	}

	w.Write(bytes)
}
