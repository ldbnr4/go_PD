package main

import "net/http"

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
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"AlbumCreate",
		"POST",
		"/album",
		AlbumCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"UserCreate",
		"POST",
		"/user",
		UserCreate,
	},
	Route{
		"PhotoCreate",
		"POST",
		"/photo",
		PhotoCreate,
	},
	Route{
		"PhotoDelete",
		"POST",
		"/del/photo",
		PhotoDelete,
	},
	Route{
		"UserDelete",
		"POST",
		"/del/user",
		UserDelete,
	},
	Route{
		"GetAlbum",
		"GET",
		"/album",
		GetAlbum,
	},
	Route{
		"GetPhoto",
		"GET",
		"/photo",
		GetPhoto,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
}
