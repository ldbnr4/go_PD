package main

// UserRoutes ...
var UserRoutes = Routes{
	Route{
		"ProfPic",
		"GET",
		"/user/:UID",
		ProfPic,
	},
	Route{
		"GetFriends",
		"GET",
		"/friends/:UID",
		GetFriends,
	},
	Route{
		"SearchUser",
		"GET",
		"/user/s/:NAME_LIKE",
		SearchUser,
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
		"Login",
		"POST",
		"/login",
		Login,
	},
}
