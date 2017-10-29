package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range getAllRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

func getAllRoutes() []Route {
	routesArray := []Routes{
		routes,
		AlbumRoutes,
		UserRoutes,
		PhotoRoutes,
	}

	var retRoutes []Route
	for _, array := range routesArray {
		for _, route := range array {
			retRoutes = append(retRoutes, route)
		}
	}

	return retRoutes
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
