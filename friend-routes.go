package main

// AlbumRoutes ...
var FriendRoutes = Routes{
	Route{
		"NetworkRelation",
		"GET",
		"/friend/relation/:nickname",
		networkRelation,
	},
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
		"POST",
		"/friend/disconnect",
		DisconnectFriend,
	},
}
