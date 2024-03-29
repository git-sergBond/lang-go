// Package fetch
// Fetch data from url
package fetch

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func FetchV2() {
	// read command line arguments
	urlsFlag := flag.String("urls", "", "www.google.ru,https://dribbble.com/tags/simple-website")
	flag.Parse()

	// parse -urls flag
	urls := strings.Split(*urlsFlag, ",")
	log.Printf("urls: %v", urls)

	// validation http
	for i, url := range urls {
		hasHttp := strings.HasPrefix(url, "http://")
		hasHttps := strings.HasPrefix(url, "https://")

		if !(hasHttp || hasHttps) {
			urls[i] = "https://" + url
		}
	}

	countUrls := len(urls)
	for i, url := range urls {
		// log info about request
		log.Println()
		log.Printf("url[%d/%d]: %s", i+1, countUrls, url)

		// fetch http resource by URL
		response, respErr := http.Get(url)
		if respErr != nil {
			log.Fatalf("fetch/http.Get: %s %v", url, respErr)
		}

		// print status
		log.Println()
		log.Printf("STATUS: %s", response.Status)

		// print body
		_, copyError := io.Copy(os.Stdout, response.Body)
		if copyError != nil {
			log.Fatalf("fetch/io.Copy: %v", copyError)
		}

		// close resource
		closeErr := response.Body.Close()
		if closeErr != nil {
			log.Fatalf("fetch/response.Body.Close: %v", closeErr)
		}
	}
}
