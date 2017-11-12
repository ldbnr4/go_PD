package main

import (
	"log"
	"net/http"
	"time"

	"goji.io/pat"

	"goji.io"
)

func NewMux() *goji.Mux {
	mux2 := goji.NewMux()
	for _, route := range getAllRoutes() {
		handler := route.HandlerFunc
		handler = logger(handler, route.Name)

		switch route.Method {
		case "GET":
			{
				mux2.HandleFunc(pat.Get(route.Pattern), handler)
			}
			break
		case "POST":
			{
				mux2.HandleFunc(pat.Post(route.Pattern), handler)
			}
		case "PUT":
			{
				mux2.HandleFunc(pat.Put(route.Pattern), handler)
			}
			break
		}
		mux2.HandleFunc(pat.Options(route.Pattern), http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Headers", "UID")
				w.Header().Add("Access-Control-Allow-Headers", "ENV")
				w.Header().Set("Access-Control-Allow-Methods", "PUT")
				w.Header().Add("Access-Control-Allow-Methods", "DELETE")
			},
		))

	}
	return mux2
}

func getAllRoutes() []Route {
	routesArray := []Routes{
		routes,
		AlbumRoutes,
		UserRoutes,
		PhotoRoutes,
		FriendRoutes,
	}

	var retRoutes []Route
	for _, array := range routesArray {
		for _, route := range array {
			retRoutes = append(retRoutes, route)
		}
	}

	return retRoutes
}

func logger(inner http.Handler, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

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
