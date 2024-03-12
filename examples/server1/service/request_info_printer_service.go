package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrintRequestInfo(writer io.Writer, request *http.Request) {
	// API info
	_, err := fmt.Fprintf(writer, "URL.Path = %q", request.URL.Path)
	if err != nil {
		printRequestInfoError("request.URL.Path", request)
	}
	_, err = fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	if err != nil {
		printRequestInfoError("request.Method, request.URL, request.Proto", request)
	}
	// headers
	for header, value := range request.Header {
		_, err := fmt.Fprintf(writer, "header[%q] = %q\n", header, value)
		if err != nil {
			printRequestInfoError("request.Header[header], request.Header[value]", request)
		}
	}
	_, err = fmt.Fprintf(writer, "Host = %q\nRemoteAddr = %q\n", request.Host, request.RemoteAddr)
	if err != nil {
		printRequestInfoError("request.Host, request.RemoteAddr", request)
	}
	// print form data
	if err := request.ParseForm(); err != nil {
		log.Printf("ERROR: in request.ParseForm() err = %v", err)
	}
	for field, value := range request.Form {
		_, err := fmt.Fprintf(writer, "Form[%q] = %q\n", field, value)
		if err != nil {
			printRequestInfoError("request.Form[key], request.Form[value]", request)
		}
	}
}

func printRequestInfoError(arguments string, request *http.Request) {
	log.Printf("ERROR: in fmt.Fprintf(%s) with request = %v", arguments, request)
}
