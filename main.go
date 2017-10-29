package main

import (
	"log"
	"net/http"
)

const PrjDir = "/home/fredfreeman/Apps/PD/"

// TODO: Make all handlers inheret from one master handler
// TODO: Create test, prod, dev mongo dbs
func main() {

	router := NewRouter()
	defer log.Fatal(http.ListenAndServe(":2500", router))
}
