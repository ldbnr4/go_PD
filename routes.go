package main

import (
	"fmt"
	"net/http"
)

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
		"AlbumCreate",
		"POST",
		"/album",
		AlbumCreate,
	},
	Route{
		"AlbumDelete",
		"POST",
		"/del/album",
		AlbumDelete,
	},
	Route{
		"GetAlbum",
		"GET",
		"/album",
		GetAlbum,
	},
	Route{
		"GetAlbums",
		"GET",
		"/albums/{UID}",
		GetAlbums,
	},
	Route{
		"ProfPic",
		"GET",
		"/user/{UID}",
		ProfPic,
	},
	Route{
		"UserCreate",
		"POST",
		"/user",
		UserCreate,
	},
	Route{
		"UserDelete",
		"POST",
		"/del/user",
		UserDelete,
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
		"GetPhoto",
		"GET",
		"/photo",
		GetPhoto,
	},
	Route{
		"DevHero",
		"GET",
		"/dev/hero/{hero}",
		DevHero,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"GetFriends",
		"GET",
		"/friends/{UID}",
		GetFriends,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
