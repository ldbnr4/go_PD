package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (ctrl *PDUMgoController) InsertAlbum(title string) string {
	if title == "" {
		panic("empty title")
	}

	newAlbum := Album{Title: title, HostID: ctrl.User.ObjectId, Creation: time.Now().UTC()}
	newAlbum.ObjectId = bson.NewObjectId()

	ifErr(ctrl.albumCol.Insert(newAlbum))

	mgoAddToSet(ctrl.userCol, ctrl.User.ObjectId, "albums", newAlbum.ObjectId)

	return newAlbum.ObjectId.Hex()
}

//RemoveAlbum removes an album from the DB
func (ctrl *PDUMgoController) RemoveAlbum(aidStr string) {
	if aidStr == "" {
		panic("empty album id")
	}
	aid := bson.ObjectIdHex(aidStr)
	ctrl.RemoveAlbumInternal(aid)
}

func (ctrl *PDUMgoController) RemoveAlbumInternal(aid bson.ObjectId) {
	album := getAlbumObj(aid, ctrl.albumCol)

	mgoRmFrmSet(ctrl.userCol, ctrl.User.ObjectId, "albums", aid)

	ctrl.DeletePhotosFrmAlbum(album)

	ifErr(ctrl.albumCol.RemoveId(aid))
}

//GetAlbumsMgo ...
func (ctrl *PDUMgoController) GetAlbumsMgo() GetAlbumsResp {

	// taggedList := foundUser.Tagged
	var created []GetAlbumResp
	for _, albumID := range ctrl.User.Albums {
		album := getAlbumObj(albumID, ctrl.albumCol)
		albumResp := GetAlbumResp{
			album.Title,
			albumID.Hex(),
			ctrl.GetGuestListNickname(album),
			nil, //TODO Fill this
			album.Creation}
		created = append(created, albumResp)
	}

	// tagged := make([]GetAlbumResp, len(taggedList))
	// for i, albumID := range taggedList {
	// 	ifErr(ctrl.albumCol.Find(bson.M{"_id": albumID}).One(album))
	// 	tagged[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	// }

	// TODO Fill this
	return GetAlbumsResp{}
}

func (ctrl *PDUMgoController) GetAlbumPhotos(aidStr string) []string {
	if aidStr == "" {
		panic("Empty AlbumId in GetAlbumPhotos")
	}
	uid := ctrl.User.ObjectId

	albumObj := getAlbumObj(bson.ObjectIdHex(aidStr), ctrl.albumCol)
	if !onTheGuestList(albumObj, uid) && albumObj.HostID != uid {
		panic(fmt.Sprintf("No access rights for the UID: %s to the album %s", uid, albumObj.Title))
	}

	var pids []string

	for _, pid := range albumObj.PhotoList {
		pids = append(pids, pid.Hex())
	}

	return pids
}

func (ctrl *PDUMgoController) GetGuestListNickname(album Album) []string {
	var guestNames []string
	for _, guestUID := range album.GuestList {
		guestName := getUserObj(guestUID, ctrl.userCol).Nickname
		guestNames = append(guestNames, guestName)
	}
	return guestNames
}
