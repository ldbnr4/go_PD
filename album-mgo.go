package main

import (
	//	"fmt"
	"time"

	"gopkg.in/mgo.v2"
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
func InsertAlbum(msg AddAlbumMsg) string {
	if msg.Title == "" {
		panic("empty title")
	}
	newAlbum := newAlbum(msg)

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")

	albumsC := db.C("albums")
	ifErr(albumsC.Insert(newAlbum))

	mgoAddToSetID(db.C("accnts"), bson.ObjectIdHex(msg.UserId), "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func RemoveAlbum(msg AlbumMsgToken) {
	switch {
	case msg.UID == "":
		panic("empty user id")
	case msg.AID == "":
		panic("empty album id")
	}

	session, err := mgo.Dial("localhost:27012")
	defer session.Close()
	ifErr(err)

	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")

	albumC := db.C("albums")
	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(albumC.FindId(albumObjID).One(&albumObj))

	mgoRmFrmSetID(db.C("accnts"), bson.ObjectIdHex(msg.UID), "albums", albumObjID)

	DeletePhotosFrmAlbum(albumObj)

	ifErr(albumC.RemoveId(bson.ObjectIdHex(msg.AID)))
}

func GetAlbumPhotos(msg AlbumMsgToken) []string {
	switch {
	case msg.AID == "":
		panic("Empty AlbumId in GetAlbumPhotos")
	case msg.UID == "":
		panic("Empty UserId in GetAlbumPhotos")
	}

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	db := session.DB("test")

	albumC := db.C("albums")
	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(albumC.FindId(albumObjID).One(&albumObj))

	if !onTheGuestList(albumObj, msg.AID) && albumObj.Host.Hex() != msg.UID {
		panic("No access rights")
	}

	var pids []string

	for _, pid := range albumObj.Photos {
		pids = append(pids, pid.Hex())
	}

	return pids

}
