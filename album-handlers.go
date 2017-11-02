package main

import (
	"encoding/json"

	"net/http"

	"goji.io/pat"
)

func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	msg := r.PostFormValue("Title")
	resp := ctrl.InsertAlbum(msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(resp))
}

//AlbumDelete ...
func AlbumDelete(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	msg := r.PostFormValue("AID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctrl.RemoveAlbum(msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetAlbumPhotos ...
func GetAlbumPhotos(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	aid := pat.Param(r, "AID")
	pids := ctrl.GetAlbumPhotos(aid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(pids))
}

//GetAlbums ...
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	aids := ctrl.GetAlbumsMgo()
	ifErr(json.NewEncoder(w).Encode(aids))
}
