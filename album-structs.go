package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	bson.ObjectId        "_id"
	Title                string
	HostID               bson.ObjectId
	Creation             time.Time
	PhotoList, GuestList []bson.ObjectId
}

//GetAlbumsResp ...
type GetAlbumsResp struct {
	CreatedAlbums []GetAlbumResp
	TaggedAlbums  []GetAlbumResp
}

//GetAlbumResp ...
type GetAlbumResp struct {
	Title     string
	AID       string
	GuestList []string
	PhotoList []string
	Creation  time.Time
}

type AlbumIDMsg struct {
	AID string
}

type AlbumTitleMsg struct {
	Title string
}

type AlbumGuestListMsg struct {
	GuestList []string
}

type AlbumPhotoListMsg struct {
	PhotoList []string
}
