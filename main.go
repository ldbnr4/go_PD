package main

import (
	"log"
	"net/http"
)

// const PrjDir = "/home/fredfreeman/Apps/PD/"
const PrjDir = "~/fredfreeman/Apps/PD/"

// type MainHandler

// TODO:	~ Create prod, dev mongo dbs (add type to header in request maybe?)
// 			~ stress test requests with js
//			~ Use PUT and DELETE methods
func main() {
	mux := NewMux()
	defer log.Fatal(http.ListenAndServe(":2500", mux))
}
