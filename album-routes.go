package main

// AlbumRoutes ...
var AlbumRoutes = Routes{
	Route{
		"GetAlbumPhotos",
		"GET",
		"/album/photos/:AID",
		GetAlbumPhotos,
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
		"/album/del",
		AlbumDelete,
	},
}
