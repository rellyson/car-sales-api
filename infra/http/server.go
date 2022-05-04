package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rellyson/car-sales-api/application/utils"
)

func CreateHttpServer(addr string) {
	r := SetRoutes()
	l := utils.NewLogger()

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	l.Info(fmt.Sprintf("[HTTPServer] - Server is ready to accept connections at %v.", addr))
	l.Fatal(srv.ListenAndServe())
}
