package main

import (
	"log"
	"net/http"
)

const PrjDir = "/home/fredfreeman/Apps/PD/"

func main() {

	router := NewRouter()
	defer log.Fatal(http.ListenAndServe(":2500", router))
}
