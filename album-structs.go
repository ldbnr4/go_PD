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
	CreatedAlbums map[string]GetAlbumResp
	TaggedAlbums  map[string]GetAlbumResp
}

//GetAlbumResp ...
type GetAlbumResp struct {
	Title     string
	GuestList []string
	PhotoList []string
	Creation  time.Time
}

type AlbumGuestListMsg struct {
	GuestList []string
}

type AlbumPhotoListMsg struct {
	PhotoList []string
}

type InsertAlbumResp struct {
	Duplicate bool
	ID        string
}
