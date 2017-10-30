package main

// AlbumRoutes ...
var AlbumRoutes = Routes{
	Route{
		"GetAlbum",
		"GET",
		"/album",
		GetAlbum,
	},
	Route{
		"GetAlbums",
		"GET",
		"/albums/:UID",
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
