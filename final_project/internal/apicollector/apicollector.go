package apicollector

import (
	"log"
	"sync"
)

type ApiCollector interface {
	AddApiCall(path string)
	PrintApiInfo()
}

type apiCollector struct {
	info  map[string]int
	mutex *sync.RWMutex
}

func NewApiCollector() ApiCollector {
	return &apiCollector{
		info:  make(map[string]int),
		mutex: &sync.RWMutex{},
	}
}

func (a *apiCollector) AddApiCall(path string) {
	a.mutex.Lock()
	a.info[path]++
	a.mutex.Unlock()
}

func (a *apiCollector) PrintApiInfo() {
	for path, callsCnt := range a.info {
		log.Printf("Path: %s, call count: %d\n", path, callsCnt)
	}
}
