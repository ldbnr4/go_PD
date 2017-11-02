package main

import (
	"log"
	"net/http"
)

const PrjDir = "/home/fredfreeman/Apps/PD/"

// const PrjDir = "~/fredfreeman/Apps/PD/"

// type MainHandler

// TODO:	~ stress test requests with js
//				~ Remove all (Controller) funcs if not a direct handler connection
func main() {
	mux := NewMux()
	defer log.Fatal(http.ListenAndServe(":2500", mux))
}
