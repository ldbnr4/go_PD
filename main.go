package main

import (
	"log"
	"net/http"
)

func main() {

	mangoMango()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
