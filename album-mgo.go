package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (ctrl *Controller) InsertAlbum(title string) InsertAlbumResp {
	var resp InsertAlbumResp
	if title == "" {
		panic("empty title")
	}
	for _, aid := range ctrl.User.Albums {
		if getAlbumObj(aid, ctrl.albumCol).Title == title {
			resp.Duplicate = true
			return resp
		}
	}

	var objIdArray []bson.ObjectId
	newAlbum := Album{
		bson.NewObjectId(),
		title,
		ctrl.User.ObjectId,
		time.Now().UTC(),
		objIdArray,
		objIdArray}

	ifErr(ctrl.albumCol.Insert(newAlbum))

	mgoAddToSet(ctrl.userCol, ctrl.User.ObjectId, "albums", newAlbum.ObjectId)
	resp.ID = newAlbum.ObjectId.Hex()

	return resp
}

func (ctrl *Controller) RemoveAlbum(aidStr string) {
	if aidStr == "" {
		panic("empty album id")
	}
	aid := bson.ObjectIdHex(aidStr)
	ctrl.removeAlbumInternal(aid)
}

func (ctrl *Controller) removeAlbumInternal(aid bson.ObjectId) {
	album := getAlbumObj(aid, ctrl.albumCol)

	mgoRmFrmSet(ctrl.userCol, ctrl.User.ObjectId, "albums", aid)

	deletePhotosFrmAlbum(album, ctrl.MgoCollections)

	ifErr(ctrl.albumCol.RemoveId(aid))
}

//GetAlbumsMgo ...
func (ctrl *Controller) GetAlbumsMgo() GetAlbumsResp {
	return GetAlbumsResp{
		ctrl.makeAlbumResp(ctrl.User.Albums),
		ctrl.makeAlbumResp(ctrl.User.Tagged)}
}

func (ctrl *Controller) GetAlbumPhotos(aidStr string) []string {
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

func (ctrl *Controller) GetGuestListNickname(album Album) []string {
	var guestNames []string
	for _, guestUID := range album.GuestList {
		guestName := getUserObj(guestUID, ctrl.userCol).Nickname
		guestNames = append(guestNames, guestName)
	}
	return guestNames
}

func deletePhotosFrmAlbum(albumObj Album, collections MgoCollections) {
	if albumObj.HostID.Hex() == "" {
		panic("DeletePhotosFrmAlbum: host id empty")
	}
	for _, photoID := range albumObj.PhotoList {
		removePhotoFile(photoID.Hex(), albumObj.HostID.Hex())
		ifErr(collections.photoCol.RemoveId(photoID))
	}
}

func (ctrl *Controller) getPIDStrs(pids []bson.ObjectId) []string {
	var pidStrs []string
	for _, pid := range pids {
		pidStrs = append(pidStrs, pid.Hex())
	}
	return pidStrs
}

func (ctrl *Controller) makeAlbumResp(albums []bson.ObjectId) []GetAlbumResp {
	var resp []GetAlbumResp
	for _, albumID := range albums {
		album := getAlbumObj(albumID, ctrl.albumCol)
		albumResp := GetAlbumResp{
			album.Title,
			albumID.Hex(),
			ctrl.GetGuestListNickname(album),
			ctrl.GetAlbumPhotos(albumID.Hex()),
			album.Creation}
		resp = append(resp, albumResp)
	}
	return resp
}
