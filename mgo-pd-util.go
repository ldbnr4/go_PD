package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAlbumObj(aid bson.ObjectId, collection *mgo.Collection) Album {
	var album Album
	ifErr(collection.FindId(aid).One(&album))
	return album
}

// TODO FIX THIS
func getUserObj(uid bson.ObjectId, collection *mgo.Collection) ServerUser {
	// var buffer map[string]map[string]string
	var buffer *ServerUser
	ifErr(collection.FindId(uid).One(&buffer))
	return *buffer
}

func getPhotoObj(pid bson.ObjectId, collection *mgo.Collection) Photo {
	var photo Photo
	ifErr(collection.FindId(pid).One(&photo))
	return photo
}
