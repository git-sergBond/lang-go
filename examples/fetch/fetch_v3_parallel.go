// Package fetch
// Load data from url and print info about benchmark (how many times it does)
package fetch

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func FetchV3() {
	// read command line arguments
	urlsFlag := flag.String("urls", "", "-urls https://www.google.ru,https://dribbble.com/tags/simple-website")
	fileFlag := flag.Bool("file", true, "-file")
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
		go fetchTimeByUrl(url, getFileName(url), *fileFlag, ch)
	}

	for range urls {
		// get data from channel and print it
		log.Println(<-ch)
	}

	log.Printf("%.2fs elapsed\n", time.Since(startTime).Seconds())
}

func fetchTimeByUrl(url, fileName string, fileFlag bool, ch chan<- string) {
	startTime := time.Now()

	resp, errResp := http.Get(url)
	if errResp != nil {
		ch <- fmt.Sprint(errResp)
		return
	}

	body, errRead := io.ReadAll(resp.Body)
	errClose := resp.Body.Close()
	if errClose != nil {
		ch <- fmt.Sprint(errClose)
		return
	}
	if errRead != nil {
		ch <- fmt.Sprint(errRead)
		return
	}

	timeElapsedSec := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", timeElapsedSec, len(body), url)

	saveToFile(body, fileName, fileFlag)
}

func saveToFile(body []byte, fileName string, fileFlag bool) {
	if !fileFlag {
		return
	}
	create, errCreate := os.Create(fileName + ".html")
	if errCreate != nil {
		log.Fatal()
	}
	_, errWrite := create.Write(body)
	if errWrite != nil {
		return
	}
}

func getFileName(url string) string {
	// delete symbols in regex from string
	r1 := regexp.MustCompile("http://|https://")
	fileName := r1.ReplaceAllString(url, "")

	r2 := regexp.MustCompile("[.&?=_/]")
	fileName = r2.ReplaceAllString(fileName, "-")

	// make filename shorter
	lenName := len(fileName)
	if lenName > 50 {
		fileName = fileName[lenName-50-1:]
	}
	return fileName
}
