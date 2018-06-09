package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var testImagePath = "test-images/gopher.png"
var port = "8000"
var basePath = "http://localhost:" + port + "/"

func TestServerServesImage(t *testing.T) {
	err := New(":" + port)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.Get(basePath + "gopher.png")

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
	err := New(":8000")
	if err != nil {
		t.Error(err)
	}

	reader, err := os.Open(testImagePath)

	if err != nil {
		t.Error(err)
	}

	response, err := http.Post(basePath+"gopher2.png", "image/png", reader)

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != 202 {
		t.Errorf("Wrong status code %s", response.Status)
	}

}
