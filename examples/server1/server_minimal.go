// Package server1
// Minimal http echo server
package server1

import (
	"example.com/examples/server1/resource"
	"log"
	"net/http"
)

func ServerMinimal() {
	http.HandleFunc("/", resource.RootHandler)
	http.HandleFunc("/count", resource.GetCountRequestsHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
