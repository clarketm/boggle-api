package api

import (
	"log"
	"net/http"
)

func NewServer(addr string, errorLog *log.Logger, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  handler,
	}
}
