package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (c *MgoController) InsertPhoto(msg AddPhotoMsg, id bson.ObjectId) {
	if msg.Album == "" {
		panic("empty album")
	}

	photObj := Photo{ObjectId: id, Upload: time.Now().UTC(), Owner: c.UID, Album: bson.ObjectIdHex(msg.Album)}

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
	if delPicMsg.PID == "" {
		panic("DeletePhoto: empty pic id")
	}

	PIDObj := bson.ObjectIdHex(delPicMsg.PID)
	picObj := new(Photo)
	c.photoCol.FindId(PIDObj).One(picObj)

	albumObj := new(Album)
	c.albumCol.FindId(picObj.Album).One(albumObj)

	if picObj.Owner == c.UID || albumObj.Host == c.UID {
		FSRemovePhoto(delPicMsg.PID, c.UID.Hex())
		mgoRmFrmSetID(c.albumCol, picObj.Album, "photos", picObj.ObjectId)
		ifErr(c.photoCol.RemoveId(PIDObj))
	}

}
