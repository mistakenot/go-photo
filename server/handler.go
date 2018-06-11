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

// GetAlbum returns a summary of an album if it exists.
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumName := vars["albumName"]

	if albumName == "" {
		http.NotFound(w, r)
		return
	}

	path := path.Join("test-images", albumName)

	if _, err := os.Stat(path); err != nil {
		http.NotFound(w, r)
		return
	}

	_, err := ioutil.ReadDir(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	size := 0
	names := make([]string, len(files))

	for i := 0; i < len(files); i++ {
		size += int(files[i].Size())
		names[i] = files[i].Name()
	}

	album := Album{
		Name:  albumName,
		Size:  size,
		Count: len(files),
		Files: names,
		URL:   fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)}

	if album.Count > 0 {
		album.Thumbnail = fmt.Sprintf("http://%s%s/%s", r.Host, r.URL.Path, names[0])
	}

	bytes, err := json.Marshal(album)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(bytes)
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

		albumFileNames := make([]string, len(photos))

		for y := 0; y < len(photos); y++ {
			albumFileNames[y] = photos[y].Name()
		}

		albums[i] = Album{
			Name:  files[i].Name(),
			Size:  int(files[i].Size()),
			Count: len(photos),
			Files: albumFileNames,
			URL:   fmt.Sprintf("http://%s/api/%s", r.Host, files[i].Name())}

		if albums[i].Count > 0 {
			albums[i].Thumbnail = fmt.Sprintf("http://%s/api/%s/%s", r.Host, files[i].Name(), albumFileNames[0])
		}

		size += int(files[i].Size())
	}

	model := AlbumOverview{
		Albums: albums}

	bytes, err := json.Marshal(model)

	if err != nil {

	}

	w.Write(bytes)
}
