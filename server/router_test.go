package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

var testImagePath = "test-images/test-album/gopher.png"
var port = "8000"

func TestServerServesExistingImage(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	resp, err := http.Get(basePath + "/api/test-album/gopher.png")

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Http code %s", resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	file, err := ioutil.ReadFile(testImagePath)

	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(body, file) != 0 {
		t.Error("Byte arrays not equal")
	}
}

func TestServerGetNonExistentImage(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	resp, err := http.Get(basePath + "/api/test-image/gopasdfaher.png")

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Wrong status code %s", resp.Status)
	}
}

func TestServerUploadsImage(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	reader, err := os.Open(testImagePath)

	if err != nil {
		t.Error(err)
	}

	response, err := http.Post(basePath+"/api/test-album/gopher2.png", "image/png", reader)

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusCreated {
		statusBytes, _ := ioutil.ReadAll(response.Body)
		t.Errorf("Wrong status code %s status %s", response.Status, statusBytes)
	}

	path := path.Join("test-images", "test-album", "gopher2.png")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("File does not exist.")
	}

	defer os.Remove(path)
}

func TestServerGetExistingAlbum(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	resp, err := http.Get(basePath + "/api/test-album")

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Http code %s", resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	model := &Album{}

	json.Unmarshal(body, model)

	if err != nil {
		t.Error(err)
	}

	if model.Count != 1 {
		t.Errorf("Album is length %d instead of 1.", model.Count)
	}

	if model.Name != "test-album" {
		t.Errorf("Image name %s is wrong.", model.Name)
	}

	if model.URL != fmt.Sprintf("%s/api/test-album", basePath) {
		t.Errorf("Image name %s is wrong.", model.URL)
	}

	if model.Thumbnail != fmt.Sprintf("%s/api/test-album/gopher.png", basePath) {
		t.Errorf("Image thumbnail %s is wrong.", model.Thumbnail)
	}
}

func TestServerGetNonExistentAlbum(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	resp, err := http.Get(basePath + "/api/test-album-sadfa")

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 404 {
		t.Errorf("Http code %s", resp.Status)
	}

	defer resp.Body.Close()
}

func TestServerListCollections(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	response, err := http.Get(basePath + "/api")
	responseBytes, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		t.Error(response.Status)
	}

	if err != nil {
		t.Error(err)
	}

	model := &AlbumOverview{}

	json.Unmarshal(responseBytes, model)

	if err != nil {
		t.Error(err)
	}

	if len(model.Albums) != 1 {
		t.Errorf("Album is length %d instead of 1.", len(model.Albums))
	}

	album := model.Albums[0]

	if album.Name != "test-album" {
		t.Errorf("File has wrong name %s", album.Name)
	}

	if album.Count != 1 {
		t.Errorf("Album has wrong count %d.", album.Count)
	}

	if album.Size == 0 {
		t.Errorf("Album has wrong size %d.", album.Size)
	}

	if album.Files[0] != "gopher.png" {
		t.Errorf("Album has wrong filename %s", album.Files[0])
	}

}
