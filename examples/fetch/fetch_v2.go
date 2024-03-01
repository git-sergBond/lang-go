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
	//read command line arguments
	urlsFlag := flag.String("urls", "", "www.google.ru,https://dribbble.com/tags/simple-website")
	flag.Parse()

	//parse -urls flag
	urls := strings.Split(*urlsFlag, ",")
	log.Printf("urls: %v", urls)

	//validation http
	for i, url := range urls {
		http := strings.HasPrefix(url, "http://")
		https := strings.HasPrefix(url, "https://")

		if !(http || https) {
			urls[i] = "https://" + url
		}
	}

	countUrls := len(urls)
	for i, url := range urls {
		log.Println()
		log.Printf("url[%d/%d]: %s", i+1, countUrls, url)

		// fetch http resource by URL
		response, respErr := http.Get(url)
		if respErr != nil {
			log.Fatalf("fetch/http.Get: %s %v", url, respErr)
		}

		// print body
		_, copyError := io.Copy(os.Stdout, response.Body)
		if copyError != nil {
			log.Fatalf("fetch/response.Body.Close: %v", copyError)
		}

		closeErr := response.Body.Close()
		if closeErr != nil {
			log.Fatalf("fetch/response.Body.Close: %v", closeErr)
		}
	}
}
