package main

import (
	"github.com/toffguy77/statusPage/middleware/httpServer"
	"log"
)

func main() {
	srv, err := httpServer.NewServer("127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(srv.ListenAndServe())
}
