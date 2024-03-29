// Package server1
// Minimal http echo server
package server1

import (
	"example.com/examples/lissajous"
	"example.com/examples/server1/resource"
	"example.com/examples/server1/service"
	"log"
	"net/http"
)

func ServerMinimal() {
	http.HandleFunc("/", resource.RootHandler)
	http.HandleFunc("/count", resource.GetCountRequestsHandler)
	http.HandleFunc("/gif", resource.GifExampleHandler)

	// Two ways define anonymous functions
	http.HandleFunc("/gif1", func(writer http.ResponseWriter, request *http.Request) {
		service.IncrementCounter(request)
		lissajous.LissajousGif(writer, service.GifExampleHandlerGetParameters(request))
	})

	handler := func(writer http.ResponseWriter, request *http.Request) {
		service.IncrementCounter(request)
		lissajous.LissajousGif(writer, service.GifExampleHandlerGetParameters(request))
	}
	http.HandleFunc("/gif2", handler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
