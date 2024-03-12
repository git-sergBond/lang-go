package resource

import (
	"example.com/examples/lissajous"
	"example.com/examples/server1/service"
	"fmt"
	"log"
	"net/http"
)

// Print request path
func RootHandler(response http.ResponseWriter, request *http.Request) {
	service.IncrementCounter(request)
	service.PrintRequestInfo(response, request)
}

// Print number of requests
func GetCountRequestsHandler(response http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(response, "%d", service.GetTotalNumberHandledRequests())
	if err != nil {
		log.Printf("ERROR: in countHandler with countRequests = %d", service.GetTotalNumberHandledRequests())
	}
}

// Gif example in browser
func GifExampleHandler(response http.ResponseWriter, request *http.Request) {
	service.IncrementCounter(request)
	urlParameters := service.GifExampleHandlerGetParameters(request)
	lissajous.LissajousGif(response, urlParameters)
}
