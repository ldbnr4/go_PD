package main

// UserRoutes ...
var UserRoutes = Routes{
	Route{
		"GetFriends",
		"GET",
		"/friends/:UID",
		GetFriends,
	},
	Route{
		"SearchUser",
		"GET",
		"/user/search/:NAME_LIKE",
		SearchUser,
	},
	Route{
		"UserCreate",
		"PUT",
		"/user",
		UserCreate,
	},
	Route{
		"UserDelete",
		"DELETE",
		"/user",
		UserDelete,
	},
	Route{
		"Login",
		"GET",
		"/user",
		Login,
	},
}
