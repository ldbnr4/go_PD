package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (pdMgo *PDUIDMgoController) InsertAlbum(title string) string {
	if title == "" {
		panic("empty title")
	}

	newAlbum := Album{Title: title, HostID: pdMgo.UID, Creation: time.Now().UTC()}
	newAlbum.ObjectId = bson.NewObjectId()

	ifErr(pdMgo.albumCol.Insert(newAlbum))

	mgoAddToSet(pdMgo.userCol, pdMgo.UID, "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func (ctrl *PDUIDMgoController) RemoveAlbum(aidStr string) {
	if aidStr == "" {
		panic("empty album id")
	}

	aid := bson.ObjectIdHex(aidStr)

	album := getAlbumObj(aid, ctrl.albumCol)

	mgoRmFrmSet(ctrl.userCol, ctrl.UID, "albums", aid)

	ctrl.DeletePhotosFrmAlbum(album)

	ifErr(ctrl.albumCol.RemoveId(aid))
}

//GetAlbumsMgo ...
func (ctrl *PDUIDMgoController) GetAlbumsMgo() GetAlbumsResp {
	foundUser := getUserObj(ctrl.UID, ctrl.userCol)

	taggedList := foundUser.Tagged
	var created []GetAlbumResp
	for i, albumID := range foundUser.Albums {
		album := getAlbumObj(albumID, ctrl.albumCol)
		albumResp := GetAlbumResp{album.Title, albumID.Hex(), ctrl.GetGuestListAlbum(album), nil, album.Creation}
		created = append(created, albumResp)
	}

	// tagged := make([]GetAlbumResp, len(taggedList))
	// for i, albumID := range taggedList {
	// 	ifErr(ctrl.albumCol.Find(bson.M{"_id": albumID}).One(album))
	// 	tagged[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	// }

	return GetAlbumsResp{created}
}

func (pdMgo *PDUIDMgoController) GetAlbumPhotos(AID string) []string {
	if AID == "" {
		panic("Empty AlbumId in GetAlbumPhotos")
	}

	var albumObj Album
	ifErr(pdMgo.albumCol.FindId(bson.ObjectIdHex(AID)).One(&albumObj))
	if !onTheGuestList(albumObj, pdMgo.UID) && albumObj.HostID != pdMgo.UID {
		panic(fmt.Sprintf("No access rights for the UID: %s to the album %s", pdMgo.UID, albumObj.Title))
	}

	var pids []string

	for _, pid := range albumObj.PhotoList {
		pids = append(pids, pid.Hex())
	}

	return pids
}

func (ctrl *PDUIDMgoController) GetGuestListNickname(album Album) []string {
	var guestNames []string
	for _, guestUID := range album.GuestList {
		var guestName string
		ctrl.userCol.FindId(guestUID).Select(bson.M{"nickname": 1}).One(guestName)
		guestNames = append(guestNames, guestName)
	}
	return guestNames
}
