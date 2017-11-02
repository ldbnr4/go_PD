package main

// AlbumRoutes ...
var FriendRoutes = Routes{
	Route{
		"ConnectFriend",
		"POST",
		"/friend/connect",
		ConnectFriend,
	},
	Route{
		"AcceptFriend",
		"POST",
		"/friend/acpt",
		AcceptFriend,
	},
	Route{
		"DeclineFriend",
		"POST",
		"/friend/decl",
		DeclineFriend,
	},
	Route{
		"DisconnectFriend",
		"DELTE",
		"/friend/disconnect",
		DisconnectFriend,
	},
}
