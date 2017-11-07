package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//InsertAlbum inserts a new album into the DB
func (ctrl *Controller) InsertAlbum(title string) InsertAlbumResp {
	var resp InsertAlbumResp
	if title == "" {
		panic("empty title")
	}
	for _, aid := range ctrl.ServerUser.Albums {
		if getAlbumObj(aid, ctrl.albumCol).Title == title {
			resp.Duplicate = true
			return resp
		}
	}

	var objIdArray []bson.ObjectId
	newAlbum := Album{
		bson.NewObjectId(),
		title,
		ctrl.ServerUser.ObjectId,
		time.Now().UTC(),
		objIdArray,
		objIdArray}

	ifErr(ctrl.albumCol.Insert(newAlbum))

	mgoAddToSet(ctrl.userCol, ctrl.ServerUser.ObjectId, "albums", newAlbum.ObjectId)
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

	mgoRmFrmSet(ctrl.userCol, ctrl.ServerUser.ObjectId, "albums", aid)

	deletePhotosFrmAlbum(album, ctrl.MgoCollections)

	ifErr(ctrl.albumCol.RemoveId(aid))
}

func getAlbumsRespInternal(user ServerUser, albumCol *mgo.Collection) GetAlbumsResp {
	return makeAlbumsRespInternal(user, albumCol)
}

func makeAlbumsRespInternal(user ServerUser, albumCol *mgo.Collection) GetAlbumsResp {

	return GetAlbumsResp{
		CreatedAlbums: makeAlbumRespInternal(user.ObjectId, user.Albums, albumCol),
		TaggedAlbums:  makeAlbumRespInternal(user.ObjectId, user.Tagged, albumCol),
	}
}

//GetAlbumsMgo ...
func (ctrl *Controller) GetAlbumsMgo() GetAlbumsResp {
	return GetAlbumsResp{
		ctrl.makeAlbumResp(ctrl.ServerUser.Albums),
		ctrl.makeAlbumResp(ctrl.ServerUser.Tagged)}
}

func (ctrl *Controller) GetAlbumPhotos(aidStr string) []string {
	if aidStr == "" {
		panic("Empty AlbumId in GetAlbumPhotos")
	}

	uid := ctrl.ServerUser.ObjectId

	return getAlbumPhotosInternal(uid, ctrl.albumCol, aidStr)

}

func getAlbumPhotosInternal(uid bson.ObjectId, albumCol *mgo.Collection, aidStr string) []string {
	albumObj := getAlbumObj(bson.ObjectIdHex(aidStr), albumCol)
	if !onTheGuestList(albumObj, uid) && albumObj.HostID != uid {
		panic(fmt.Sprintf("No access rights for the UID: %s to the album %s", uid, albumObj.Title))
	}

	var pids []string

	for _, pid := range albumObj.PhotoList {
		pids = append(pids, pid.Hex())
	}

	return pids

}

// GetGuestListNickname ...
func (ctrl *Controller) GetGuestListNickname(album Album) []string {
	return getGuestListNicknameInternal(album, ctrl.userCol)
}

func getGuestListNicknameInternal(album Album, userCol *mgo.Collection) []string {
	var guestNames []string
	for _, guestUID := range album.GuestList {
		guestName := getUserObj(guestUID, userCol).Nickname
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

func (ctrl *Controller) makeAlbumResp(albums []bson.ObjectId) map[string]GetAlbumResp {
	return makeAlbumRespInternal(ctrl.ObjectId, albums, ctrl.albumCol)
}

func makeAlbumRespInternal(uid bson.ObjectId, albums []bson.ObjectId, albumCol *mgo.Collection) map[string]GetAlbumResp {
	resp := make(map[string]GetAlbumResp)

	for _, albumID := range albums {
		album := getAlbumObj(albumID, albumCol)
		albumResp := GetAlbumResp{
			album.Title,
			getGuestListNicknameInternal(album, albumCol),
			getAlbumPhotosInternal(uid, albumCol, albumID.Hex()),
			album.Creation}
		resp[albumID.Hex()] = albumResp
	}
	return resp
}
