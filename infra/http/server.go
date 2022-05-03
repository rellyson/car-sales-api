package http

import (
	"log"
	"net/http"
	"time"
)

func CreateHttpServer(addr string) {
	r := SetRoutes()

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("[HTTPServer] - Server is ready to accept connections at %v.", addr)
	log.Fatal(srv.ListenAndServe())
}
