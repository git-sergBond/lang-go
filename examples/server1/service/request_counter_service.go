package service

import (
	"log"
	"net/http"
	"os"
	"sync"
)

const DEBUG = true

var countRequests int64 = 0
var mutexForCountRequests sync.Mutex

func IncrementCounter(request *http.Request) {
	if DEBUG {
		log.Printf("DEBUG: IncrementCounter ====================================")
		PrintRequestInfo(os.Stdout, request)
	}
	mutexForCountRequests.Lock()
	countRequests++
	log.Printf("INFO: countRequests = %d\n", countRequests)
	mutexForCountRequests.Unlock()
}

func GetTotalNumberHandledRequests() int64 {
	return countRequests
}
