package main

import (
	"gosrv/src"
	"log"
	"net/http"
)

func main() {
	for pattern, handler := range gosrv.Router {
		http.HandleFunc(pattern, handler)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
