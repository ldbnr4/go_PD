package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (pdMgo *PDMgoController) InsertAlbum(title string) string {
	if title == "" {
		panic("empty title")
	}

	newAlbum := Album{Title: title, Host: pdMgo.UID, Creation: time.Now().UTC()}
	newAlbum.ObjectId = bson.NewObjectId()

	ifErr(pdMgo.albumCol.Insert(newAlbum))

	mgoAddToSetID(pdMgo.userCol, pdMgo.UID, "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func (pdMgo *PDMgoController) RemoveAlbum(AID string) {
	if msg.AID == "" {
		panic("empty album id")
	}

	albumObjID := bson.ObjectIdHex(msg.AID)

	var albumObj Album

	ifErr(pdMgo.albumCol.FindId(albumObjID).One(&albumObj))

	mgoRmFrmSetID(pdMgo.userCol, pdMgo.UID, "albums", albumObjID)

	pdMgo.DeletePhotosFrmAlbum(albumObj)

	ifErr(pdMgo.albumCol.RemoveId(bson.ObjectIdHex(msg.AID)))
}

func (pdMgo *PDMgoController) GetAlbumPhotos(msg AlbumMsgToken) []string {
	if msg.AID == "" {
		panic("Empty AlbumId in GetAlbumPhotos")
	}

	var albumObj Album
	ifErr(pdMgo.albumCol.FindId(bson.ObjectIdHex(msg.AID)).One(&albumObj))
	if !onTheGuestList(albumObj, pdMgo.UID) && albumObj.Host != pdMgo.UID {
		panic(fmt.Sprintf("No access rights for the UID: %s to the album %s", pdMgo.UID, albumObj.Title))
	}

	var pids []string

	for _, pid := range albumObj.Photos {
		pids = append(pids, pid.Hex())
	}

	return pids

}
