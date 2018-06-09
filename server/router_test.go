package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

var testImagePath = "test-images/gopher.png"
var port = "8000"

func TestServerServesImage(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	resp, err := http.Get(basePath + "/gopher.png")

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

func TestServerUploadsImage(t *testing.T) {
	router := New()
	server := httptest.NewServer(router)
	basePath := server.URL

	defer server.Close()

	reader, err := os.Open(testImagePath)

	if err != nil {
		t.Error(err)
	}

	response, err := http.Post(basePath+"/gopher2.png", "image/png", reader)

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusCreated {
		statusBytes, _ := ioutil.ReadAll(response.Body)
		t.Errorf("Wrong status code %s status %s", response.Status, statusBytes)
	}

	path := path.Join("test-images", "gopher2.png")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("File does not exist.")
	}

	defer os.Remove(path)
}
