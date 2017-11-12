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
		"/photo/:PID/:UID/:ENV",
		GetPhoto,
	},
	Route{
		"DevHero",
		"GET",
		"/photo/hero/:hero",
		DevHero,
	},
	Route{
		"ProfPic",
		"GET",
		"/photo/prof/:NICKNAME/:UID/:ENV",
		profPic,
	},
}
