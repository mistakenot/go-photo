package main

import (
	"net/http"

	"github.com/mistakenot/gopic/server"
)

func main() {
	router := server.New()
	http.ListenAndServe(":8000", router)
}
