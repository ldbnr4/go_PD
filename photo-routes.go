package main

var PhotoRoutes = []Route{
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
}
