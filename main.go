package main

import (
	"log"
	"net/http"

	"github.com/georgevazj/gomonitor/ansible"
	"github.com/georgevazj/gomonitor/dns"
	"github.com/georgevazj/gomonitor/iaas"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dns", dns.GetDNS).Methods("GET")
	router.HandleFunc("/ansible", ansible.GetAnsible).Methods("GET")
	router.HandleFunc("/iaas", iaas.GetIaas).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
