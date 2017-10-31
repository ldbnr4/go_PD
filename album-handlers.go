package main

import (
	"encoding/json"
	//"fmt"

	"net/http"
)

func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	msg := new(AddAlbumMsg)

	FillStruct(r, msg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getController(r)
	defer ctrl.session.Close()
	aid := ctrl.InsertAlbum(*msg)
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(&AddAlbumResp{AID: aid, Title: msg.Title}))
}

//AlbumDelete ...
func AlbumDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(AlbumMsgToken)
	FillStruct(r, msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.RemoveAlbum(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetAlbum ...
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	msg := AlbumMsgToken{AID: r.URL.Query().Get("AlbumId")}
	pids := ctrl.GetAlbumPhotos(msg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(&GetAlbumPIDsResp{PhotoIDs: pids}))
}
