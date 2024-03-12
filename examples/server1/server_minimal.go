// Package server1
// Minimal http echo server
package server1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var countRequests int64 = 0
var mutexForCountRequests sync.Mutex

func ServerMinimal() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Print request path
func handler(response http.ResponseWriter, request *http.Request) {
	incrementCounter(request)
	_, err := fmt.Fprintf(response, "URL.Path = %q", request.URL.Path)
	if err != nil {
		log.Printf("ERROR: in handler/Fprintf with request.URL.Path argument = " + request.URL.Path)
	}
}

// Print number of requests
func countHandler(response http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(response, "%d", countRequests)
	if err != nil {
		log.Printf("ERROR: in handler/countHandler with countRequests = %d", countRequests)
	}
}

func incrementCounter(request *http.Request) {
	log.Printf("DEBUG: reqest = %v", request)
	mutexForCountRequests.Lock()
	countRequests++
	log.Printf("INFO: countRequests = %d", countRequests)
	mutexForCountRequests.Unlock()
}
