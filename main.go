package main

import (
	"log"
	"net/http"
)

// const PrjDir = "/home/fredfreeman/Apps/PD/"
const PrjDir = "~/fredfreeman/Apps/PD/"

// type MainHandler

// TODO:	~ Create prod, dev mongo dbs (add type to header in request maybe?)
// 				~ stress test requests with js
//				~ Use PUT and DELETE methods
// 				~ Remove all FillStruct and replace with r.PostFormValue("key") for single value posts
func main() {
	mux := NewMux()
	defer log.Fatal(http.ListenAndServe(":2500", mux))
}
