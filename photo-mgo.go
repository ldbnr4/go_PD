package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Photo struct {
	bson.ObjectId "_id"
	Upload        time.Time
	Owner         bson.ObjectId
	Album         bson.ObjectId
}

func newPhoto(owner, album string, id bson.ObjectId) Photo {
	return Photo{id, time.Now().UTC(), bson.ObjectIdHex(owner), bson.ObjectIdHex(album)}
}

func (c *MgoController) InsertPhoto(msg AddPhotoMsg, id bson.ObjectId) {
	switch {
	case msg.Owner == "":
		panic("empty owner")
	case msg.Album == "":
		panic("empty album")
	default:
		break
	}

	photObj := newPhoto(msg.Owner, msg.Album, id)

	ifErr(c.photoCol.Insert(photObj))

	mgoAddToSetID(c.albumCol, bson.ObjectIdHex(msg.Album), "photos", id)
}

func (c *MgoController) DeletePhotosFrmAlbum(albumObj Album) {
	switch {
	case albumObj.Host.Hex() == "":
		panic("DeletePhotosFrmAlbum: empty user id")
	}
	for _, photoID := range albumObj.Photos {
		FSRemovePhoto(photoID.Hex(), albumObj.Host.Hex())
		ifErr(c.photoCol.RemoveId(photoID))
	}
}

func (c *MgoController) DeletePhoto(delPicMsg DelPhotoMsg) {
	switch {
	case delPicMsg.PID == "":
		panic("DeletePhoto: empty pic id")
	case delPicMsg.UID == "":
		panic("DeletePhoto: empty user id")
	}

	PIDObj := bson.ObjectIdHex(delPicMsg.PID)
	picObj := new(Photo)
	c.photoCol.FindId(PIDObj).One(picObj)

	albumObj := new(Album)
	c.albumCol.FindId(picObj.Album).One(albumObj)

	if picObj.Owner.Hex() == delPicMsg.UID || albumObj.Host.Hex() == delPicMsg.UID {
		FSRemovePhoto(delPicMsg.PID, delPicMsg.UID)
		mgoRmFrmSetID(c.albumCol, picObj.Album, "photos", picObj.ObjectId)
		ifErr(c.photoCol.RemoveId(PIDObj))
	}

}
