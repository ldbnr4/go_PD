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
	bson.ObjectId        "_id"
	Title                string
	Host                 bson.ObjectId
	Creation             time.Time
	Photos, Tags, Guests []bson.ObjectId
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
	case msg.UserId == "":
		panic("empty user id")
	case msg.AlbumId == "":
		panic("empty album id")
	}

	session, err := mgo.Dial("localhost:27012")
	defer session.Close()
	ifErr(err)

	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")

	albumC := db.C("albums")
	albumObjID := bson.ObjectIdHex(msg.AlbumId)

	var albumObj Album

	ifErr(albumC.FindId(albumObjID).One(&albumObj))

	mgoRmFrmSetID(db.C("accnts"), bson.ObjectIdHex(msg.UserId), "albums", albumObjID)

	DeletePhotosFrmAlbum(albumObj)

	ifErr(albumC.RemoveId(bson.ObjectIdHex(msg.AlbumId)))
}

func GetAlbumPhotos(msg AlbumMsgToken) []string {
	switch {
	case msg.AlbumId == "":
		panic("Empty AlbumId in GetAlbumPhotos")
	case msg.UserId == "":
		panic("Empty UserId in GetAlbumPhotos")
	}

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	db := session.DB("test")

	albumC := db.C("albums")
	albumObjID := bson.ObjectIdHex(msg.AlbumId)

	var albumObj Album

	ifErr(albumC.FindId(albumObjID).One(&albumObj))

	var allowed bool

	for _, guest := range albumObj.Guests {
		if guest.Hex() == msg.AlbumId {
			allowed = true
			break
		}
	}

	if !allowed && albumObj.Host.Hex() != msg.UserId {
		panic("No access rights")
	}

	var pids []string

	for _, pid := range albumObj.Photos {
		pids = append(pids, pid.Hex())
	}

	return pids

}
