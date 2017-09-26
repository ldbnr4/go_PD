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

type GetAlbumResp struct {
	PhotoIDs []string
}

type AlbumMsgToken struct {
	UserId  string
	AlbumId string
}

/************
 * HANDLERS *
 ************/
func AlbumCreate(w http.ResponseWriter, r *http.Request) {
	msg := new(AddAlbumMsg)

	FillStruct(r, msg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	aid := InsertAlbum(*msg)
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(&AddAlbumResp{ID: aid, Title: msg.Title}))
}

func AlbumDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(AlbumMsgToken)
	FillStruct(r, msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	RemoveAlbum(*msg)

}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	msg := new(AlbumMsgToken)

	msg.AlbumId = r.URL.Query().Get("AlbumId")
	msg.UserId = r.URL.Query().Get("UserId")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pids := GetAlbumPhotos(*msg)
	ifErr(json.NewEncoder(w).Encode(&GetAlbumResp{PhotoIDs: pids}))
}
