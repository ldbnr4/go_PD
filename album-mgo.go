package main

import (
	//	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//Album that is stored in the DB
type Album struct {
	bson.ObjectId "_id"
	Title         string
	Host          bson.ObjectId
	Creation      time.Time
	Photos, Guest []bson.ObjectId
}

func newAlbum(msg AddAlbumMsg) Album {
	newU := Album{Title: msg.Title, Host: bson.ObjectIdHex(msg.UserId), Creation: time.Now().UTC()}
	newU.ObjectId = bson.NewObjectId()
	return newU
}

//InsertAlbum inserts a new album into the DB
func (c *MgoController) InsertAlbum(msg AddAlbumMsg) string {
	if msg.Title == "" {
		panic("empty title")
	}
	newAlbum := newAlbum(msg)

	ifErr(c.albumCol.Insert(newAlbum))

	mgoAddToSetID(c.userCol, bson.ObjectIdHex(msg.UserId), "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func (c *MgoController) RemoveAlbum(msg AlbumMsgToken) {
	switch {
	case msg.UID == "":
		panic("empty user id")
	case msg.AID == "":
		panic("empty album id")
	}

	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(c.albumCol.FindId(albumObjID).One(&albumObj))

	mgoRmFrmSetID(c.userCol, bson.ObjectIdHex(msg.UID), "albums", albumObjID)

	c.DeletePhotosFrmAlbum(albumObj)

	ifErr(c.albumCol.RemoveId(bson.ObjectIdHex(msg.AID)))
}

func (c *MgoController) GetAlbumPhotos(msg AlbumMsgToken) []string {
	switch {
	case msg.AID == "":
		panic("Empty AlbumId in GetAlbumPhotos")
	case msg.UID == "":
		panic("Empty UserId in GetAlbumPhotos")
	}

	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(c.albumCol.FindId(albumObjID).One(&albumObj))

	if !onTheGuestList(albumObj, msg.AID) && albumObj.Host.Hex() != msg.UID {
		panic("No access rights")
	}

	var pids []string

	for _, pid := range albumObj.Photos {
		pids = append(pids, pid.Hex())
	}

	return pids

}
