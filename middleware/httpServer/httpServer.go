package httpServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/middleware/common"
	"log"
	"net"
	"net/http"
	"time"
)

func NewServer(location string) (*http.Server, error) {
	if !correctLoc(location) {
		return nil, errors.New("server location is not valid IP[:PORT]")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         location,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv, nil
}

func correctLoc(location string) bool {
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
	fmt.Fprint(w, string(resJson))
}
