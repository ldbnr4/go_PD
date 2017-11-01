package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAlbumObj(aid bson.ObjectId, collection *mgo.Collection) Album {
	var album Album
	ifErr(collection.FindId(aid).One(album))
	return album
}

func getUserObj(uid bson.ObjectId, collection *mgo.Collection) User {
	var user User
	ifErr(collection.FindId(uid).One(user))
	return user
}

func getPhotoObj(pid bson.ObjectId, collection *mgo.Collection) Photo {
	var photo Photo
	ifErr(collection.FindId(pid).One(photo))
	return photo
}
