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
		"PUT",
		"/album",
		AlbumCreate,
	},
	Route{
		"AlbumDelete",
		"DELTE",
		"/album",
		AlbumDelete,
	},
}
