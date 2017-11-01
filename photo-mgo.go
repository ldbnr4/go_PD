package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (c *PDUIDMgoController) InsertPhoto(albumID, photoID bson.ObjectId) {
	if msg.Album == "" {
		panic("empty album")
	}

	photObj := Photo{ObjectId: id, Upload: time.Now().UTC(), Owner: c.UID, Album: bson.ObjectIdHex(msg.Album)}

	ifErr(c.photoCol.Insert(photObj))

	mgoAddToSet(c.albumCol, bson.ObjectIdHex(msg.Album), "photos", id)
}

func (c *PDUIDMgoController) DeletePhotosFrmAlbum(albumObj Album) {
	switch {
	case albumObj.Host.Hex() == "":
		panic("DeletePhotosFrmAlbum: empty user id")
	}
	for _, photoID := range albumObj.Photos {
		FSRemovePhoto(photoID.Hex(), albumObj.Host.Hex())
		ifErr(c.photoCol.RemoveId(photoID))
	}
}

func (c *PDUIDMgoController) DeletePhoto(pid string) {
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
		mgoRmFrmSet(c.albumCol, picObj.Album, "photos", picObj.ObjectId)
		ifErr(c.photoCol.RemoveId(PIDObj))
	}

}
