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
		"DELETE",
		"/photo",
		PhotoDelete,
	},
	Route{
		"GetPhoto",
		"GET",
		"/photo/:PID",
		GetPhoto,
	},
	Route{
		"DevHero",
		"GET",
		"/photo/hero/:hero",
		DevHero,
	},
}
