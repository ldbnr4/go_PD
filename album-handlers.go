package main

import (
	"encoding/json"
	//"fmt"

	"net/http"

	"goji.io/pat"
)

func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	var msg AlbumTitleMsg
	FillStruct(r, msg)
	ctrl := getPDUIDController(r)
	defer ctrl.session.Close()
	aid := ctrl.InsertAlbum(msg.Title)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(&AlbumIDMsg{AID: aid}))
}

//AlbumDelete ...
func AlbumDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(AlbumIDMsg)
	FillStruct(r, msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getPDUIDController(r)
	defer ctrl.session.Close()
	ctrl.RemoveAlbum(msg.AID)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetAlbumPhotos ...
func GetAlbumPhotos(w http.ResponseWriter, r *http.Request) {
	aid := pat.Param(r, "AID")
	ctrl := getPDUIDController(r)
	defer ctrl.session.Close()
	pids := ctrl.GetAlbumPhotos(aid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(pids))
}

//GetAlbums ...
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUIDController(r)
	defer ctrl.session.Close()
	aids := ctrl.GetAlbumsMgo()
	ifErr(json.NewEncoder(w).Encode(aids))
}
