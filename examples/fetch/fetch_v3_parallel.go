// Package fetch
// Load data from url and print info about benchmark (how many time it does)
package fetch

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func FetchV3() {
	// read command line arguments
	urlsFlag := flag.String("urls", "", "https://www.google.ru,https://dribbble.com/tags/simple-website")
	flag.Parse()

	// parse -urls flag
	urls := strings.Split(*urlsFlag, ",")
	if len(urls) <= 0 {
		log.Fatal("empty argument -urls")
	}

	// validation http
	for i, url := range urls {
		hasHttp := strings.HasPrefix(url, "http://")
		hasHttps := strings.HasPrefix(url, "https://")

		if !(hasHttp || hasHttps) {
			urls[i] = "https://" + url
		}
	}

	ch := make(chan string)
	startTime := time.Now()
	for _, url := range urls {
		// run parallel subprograms
		go fetchTimeByUrl(url, ch)
	}

	for range urls {
		// get data from channel and print it
		log.Println(<-ch)
	}

	log.Printf("%.2fs elapsed\n", time.Since(startTime).Seconds())
}

func fetchTimeByUrl(url string, ch chan<- string) {
	startTime := time.Now()

	resp, errResp := http.Get(url)
	if errResp != nil {
		ch <- fmt.Sprint(errResp)
		return
	}

	nBytes, errCopy := io.Copy(io.Discard, resp.Body)
	err := resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(errResp)
		return
	}
	if errCopy != nil {
		ch <- fmt.Sprint(errCopy)
		return
	}

	timeElapsedSec := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", timeElapsedSec, nBytes, url)
}
