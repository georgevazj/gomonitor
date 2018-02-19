package main

import (
	"log"
	"net/http"

	"gomonitor/dns"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dns", dns.GetDNS).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
