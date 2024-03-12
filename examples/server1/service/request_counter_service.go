package service

import (
	"log"
	"net/http"
	"sync"
)

const DEBUG = false

var countRequests int64 = 0
var mutexForCountRequests sync.Mutex

func IncrementCounter(request *http.Request) {
	if DEBUG {
		log.Printf("DEBUG: reqest = %v", request)
	}
	mutexForCountRequests.Lock()
	countRequests++
	log.Printf("INFO: countRequests = %d\n", countRequests)
	mutexForCountRequests.Unlock()
}

func GetTotalNumberHandledRequests() int64 {
	return countRequests
}
