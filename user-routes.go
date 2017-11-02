package main

// UserRoutes ...
var UserRoutes = Routes{
	Route{
		"ProfPic",
		"GET",
		"/user/pic/:UID",
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
	Route{
		"AcceptFriendRequest",
		"POST",
		"friend/acpt",
		AcceptReq,
	},
	Route{
		"DeclineFriendRequest",
		"POST",
		"friend/decl",
		DeclineReq,
	},
}
