package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	bson.ObjectId "_id"
	Title         string
	Host          bson.ObjectId
	Creation      time.Time
	Photos, Guest []bson.ObjectId
}

type AddAlbumResp struct {
	Title string
	AID   string
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

type AlbumCreateMsg struct {
	Title string
}
