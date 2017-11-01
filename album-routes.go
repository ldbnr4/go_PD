package main

// AlbumRoutes ...
var AlbumRoutes = Routes{
	Route{
		"GetAlbum",
		"GET",
		"/album/:AID",
		GetAlbum,
	},
	Route{
		"GetAlbums",
		"GET",
		"/albums",
		GetAlbums,
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
}
