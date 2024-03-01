package fetch

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
)

func Fetch() {
	urlsFlag := flag.String("urls", "", "https://www.google.ru,https://dribbble.com/tags/simple-website")
	flag.Parse()

	urls := strings.Split(*urlsFlag, ",")
	log.Printf("urls: %v", urls)

	countUrls := len(urls)
	for i, url := range urls {
		log.Printf("url[%d/%d]: %s", i+1, countUrls, url)

		// fetch http resource by URL
		response, respErr := http.Get(url)
		if respErr != nil {
			log.Fatalf("fetch/http.Get: %s %v", url, respErr)
		}

		// read body
		body, readErr := io.ReadAll(response.Body)
		closeErr := response.Body.Close()
		if readErr != nil {
			log.Fatalf("fetch/io.ReadAll: %v", readErr)
		}
		if closeErr != nil {
			log.Fatalf("fetch/response.Body.Close: %v", closeErr)
		}

		// print body
		log.Printf("%s", body)
	}
}
