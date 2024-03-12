package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrintRequestInfo(writer io.Writer, request *http.Request) {
	// API info
	if _, err := fmt.Fprintf(writer, "URL.Path = %q", request.URL.Path); err != nil {
		FprintfError("request.URL.Path", request, err)
	}
	if _, err := fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto); err != nil {
		FprintfError("request.Method, request.URL, request.Proto", request, err)
	}
	// headers
	for header, value := range request.Header {
		if _, err := fmt.Fprintf(writer, "header[%q] = %q\n", header, value); err != nil {
			FprintfError("request.Header[header], request.Header[value]", request, err)
		}
	}
	if _, err := fmt.Fprintf(writer, "Host = %q\nRemoteAddr = %q\n", request.Host, request.RemoteAddr); err != nil {
		FprintfError("request.Host, request.RemoteAddr", request, err)
	}
	// print form data
	if err := request.ParseForm(); err != nil {
		log.Printf("ERROR: in request.ParseForm() err = %v", err)
	}
	for field, value := range request.Form {
		if _, err := fmt.Fprintf(writer, "Form[%q] = %q\n", field, value); err != nil {
			FprintfError("request.Form[key], request.Form[value]", request, err)
		}
	}
}

func FprintfError(arguments string, request *http.Request, err error) {
	log.Printf("ERROR: in fmt.Fprintf(%s) with request = %v\n ERROR DETAILS: %v", arguments, request, err)
}
