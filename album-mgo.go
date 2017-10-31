package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (c *MgoController) InsertAlbum(msg AddAlbumMsg) string {
	if msg.Title == "" {
		panic("empty title")
	}

	newAlbum := Album{Title: msg.Title, Host: c.UID, Creation: time.Now().UTC()}
	newAlbum.ObjectId = bson.NewObjectId()

	ifErr(c.albumCol.Insert(newAlbum))

	mgoAddToSetID(c.userCol, c.UID, "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func (c *MgoController) RemoveAlbum(msg AlbumMsgToken) {
	if msg.AID == "" {
		panic("empty album id")
	}

	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(c.albumCol.FindId(albumObjID).One(&albumObj))

	mgoRmFrmSetID(c.userCol, c.UID, "albums", albumObjID)

	c.DeletePhotosFrmAlbum(albumObj)

	ifErr(c.albumCol.RemoveId(bson.ObjectIdHex(msg.AID)))
}

func (c *MgoController) GetAlbumPhotos(msg AlbumMsgToken) []string {
	if msg.AID == "" {
		panic("Empty AlbumId in GetAlbumPhotos")
	}

	var albumObj Album
	ifErr(c.albumCol.FindId(bson.ObjectIdHex(msg.AID)).One(&albumObj))
	if !onTheGuestList(albumObj, c.UID) && albumObj.Host != c.UID {
		panic("No access rights")
	}

	var pids []string

	for _, pid := range albumObj.Photos {
		pids = append(pids, pid.Hex())
	}

	return pids

}
