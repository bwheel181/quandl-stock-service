package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("latest/{ticker}", GetLatest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
