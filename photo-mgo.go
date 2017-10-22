package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Photo struct {
	bson.ObjectId "_id"
	Upload        time.Time
	Owner         bson.ObjectId
	Album         bson.ObjectId
}

func newPhoto(owner string, album string, id bson.ObjectId) Photo {
	return Photo{id, time.Now().UTC(), bson.ObjectIdHex(owner), bson.ObjectIdHex(album)}
}

func InsertPhoto(msg AddPhotoMsg, id bson.ObjectId) {
	switch {
	case msg.Owner == "":
		panic("empty owner")
	case msg.Album == "":
		panic("empty album")
	default:
		break
	}

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

	photObj := newPhoto(msg.Owner, msg.Album, id)

	db := session.DB("test")

	photosC := db.C("photos")
	ifErr(photosC.Insert(photObj))

	mgoAddToSetID(db.C("albums"), bson.ObjectIdHex(msg.Album), "photos", id)
}

func DeletePhotosFrmAlbum(albumObj Album) {
	switch {
	case albumObj.Host.Hex() == "":
		panic("DeletePhotosFrmAlbum: empty user id")
	}
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	db := session.DB("test")
	photosC := db.C("photos")

	for _, photoID := range albumObj.Photos {
		FSRemovePhoto(photoID.Hex(), albumObj.Host.Hex())
		ifErr(photosC.RemoveId(photoID))
	}
}

func DeletePhoto(delPicMsg DelPhotoMsg) {
	switch {
	case delPicMsg.PID == "":
		panic("DeletePhoto: empty pic id")
	case delPicMsg.UID == "":
		panic("DeletePhoto: empty user id")
	}

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	db := session.DB("test")
	photosC := db.C("photos")

	PIDObj := bson.ObjectIdHex(delPicMsg.PID)
	picObj := new(Photo)
	photosC.FindId(PIDObj).One(picObj)

	albumObj := new(Album)
	albumC := db.C("albums")
	albumC.FindId(picObj.Album).One(albumObj)

	if picObj.Owner.Hex() == delPicMsg.UID || onTheGuestList(*albumObj, delPicMsg.UID) {
		FSRemovePhoto(delPicMsg.PID, delPicMsg.UID)
		mgoRmFrmSetID(albumC, picObj.Album, "photos", picObj.ObjectId)
		ifErr(photosC.RemoveId(PIDObj))
	}

}
