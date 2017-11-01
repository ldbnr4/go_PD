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
		"/photo/del",
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
		"/photo/hero/{hero}",
		DevHero,
	},
}
