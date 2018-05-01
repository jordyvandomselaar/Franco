package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitServer(r *mux.Router) {
	s := &http.Server{
		Addr:    "localhost:8000",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}
