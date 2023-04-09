package httpServer

import (
	"encoding/json"
	"fmt"
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/middleware/common"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
	"log"
	"net"
	"net/http"
	"time"
)

func Run(location string) {
	if !isCorrectLoc(location) {
		log.Fatalf("server location is not correct: %s\n", location)
	}
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		log.Fatal(err)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(1*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := http.HandlerFunc(handleConnection)

	http.Handle("/", cacheClient.Middleware(handler))
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func isCorrectLoc(location string) bool {
	host, port, err := net.SplitHostPort(location)
	if err != nil {
		return false
	}

	if location == host || location == fmt.Sprintf("%s:%s", host, port) {
		return true
	}
	return false
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	var res models.ResultT
	results := common.GetResultData()
	if !common.CheckResults(results) {
		res.Status = false
		res.Error = "Error on collect data"
		w.WriteHeader(http.StatusFailedDependency)
	} else {
		res.Status = true
		res.Data = results
		w.WriteHeader(http.StatusOK)
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marchaling json answer: %v\n", err)
		return
	}

	_, err = fmt.Fprint(w, string(resJson))
	if err != nil {
		log.Printf("error sending http response: %v", err)
	}
}
