package main

import (
	"fmt"
	"net/http"
)

// TODO: replace HandlerFunc with a new type that will consume a ctrl
// func type example- type HandlerFunc func(ResponseWriter, *Request)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
