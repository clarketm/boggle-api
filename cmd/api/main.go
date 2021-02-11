package main

import (
	"log"
	"net/http"

	"github.com/clarketm/boggle-api/pkg/api"
)

func main() {
	s := api.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s))
}
