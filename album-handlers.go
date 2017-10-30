package main

import (
	"encoding/json"
	//"fmt"

	"net/http"
)

/***********
 * STRUCTS *
 ***********/
type AddAlbumMsg struct {
	Title  string
	UserId string
}

type AddAlbumResp struct {
	Title string
	ID    string
}

//GetAlbumPIDsResp ...
type GetAlbumPIDsResp struct {
	PhotoIDs []string
}

//GetAlbumsResp ...
type GetAlbumsResp struct {
	Created []GetAlbumResp
	Tagged  []GetAlbumResp
}

//GetAlbumResp ...
type GetAlbumResp struct {
	Title string
	ID    string
}

//AlbumMsgToken ...
type AlbumMsgToken struct {
	UID string
	AID string
}

/************
 * HANDLERS *
 ************/
func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	msg := new(AddAlbumMsg)

	FillStruct(r, msg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getController()
	defer ctrl.session.Close()
	aid := ctrl.InsertAlbum(*msg)
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(&AddAlbumResp{ID: aid, Title: msg.Title}))
}

//AlbumDelete ...
func AlbumDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(AlbumMsgToken)
	FillStruct(r, msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getController()
	defer ctrl.session.Close()
	ctrl.RemoveAlbum(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetAlbum ...
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	msg := AlbumMsgToken{UID: r.URL.Query().Get("UserId"), AID: r.URL.Query().Get("AlbumId")}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctrl := getController()
	defer ctrl.session.Close()
	pids := ctrl.GetAlbumPhotos(msg)
	ifErr(json.NewEncoder(w).Encode(&GetAlbumPIDsResp{PhotoIDs: pids}))
}
