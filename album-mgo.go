package main

import (
	//	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	//	"log"
)

type Album struct {
	bson.ObjectId 				"_id"
	AddAlbumMsg
	Creation         string
	Photos, Tags	[]string
}

func newAlbum(msg AddAlbumMsg) Album {
	newU := Album{AddAlbumMsg: msg, Creation: time.Now().UTC().String()}
	newU.ObjectId = bson.NewObjectId()
	return newU
}

func InsertAlbum(msg AddAlbumMsg) string {
	if msg.Title == ""{
		panic("empty title")
	}
	session, err := mgo.Dial("localhost:27012")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	
	newAlbum := newAlbum(msg)
	
	db := session.DB("test")
	
	albumsC := db.C("albums")
	err = albumsC.Insert(newAlbum)
	if err != nil {
		panic(err)
	}

	update := make(map[string]map[string]bson.ObjectId)
	op := make(map[string]bson.ObjectId)
	op["albums"] = newAlbum.ObjectId
	update["$addToSet"] = op

	accntsC := db.C("accnts")
	err = accntsC.UpdateId(bson.ObjectIdHex(msg.UserId), update)

	if err != nil {
		panic(err)
	}
	
	return newAlbum.ObjectId.Hex()
}
