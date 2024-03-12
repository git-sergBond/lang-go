// Package server1
// Minimal http echo server
package server1

import (
	"fmt"
	"log"
	"net/http"
)

func ServerMinimal() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "URL.Path = %q", request.URL.Path)
	if err != nil {
		log.Printf("ERROR: in handler/Fprintf with request.URL.Path argument = " + request.URL.Path)
	}
}
