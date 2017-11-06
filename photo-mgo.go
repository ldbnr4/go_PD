package main

import (
	"mime/multipart"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (ctrl *Controller) InsertPhoto(file multipart.File, aidStr string) {
	if aidStr == "" {
		panic("empty album")
	}
	aid := bson.ObjectIdHex(aidStr)
	pid := bson.NewObjectId()

	SaveImageFile(file, ctrl.User.ObjectId.Hex(), pid.Hex())

	photObj := Photo{
		pid,
		time.Now().UTC(),
		ctrl.User.ObjectId,
		aid}

	ifErr(ctrl.photoCol.Insert(photObj))

	mgoAddToSet(ctrl.albumCol, aid, "photolist", pid)
}

func (ctrl *Controller) DeletePhoto(pidStr string) {
	if pidStr == "" {
		panic("DeletePhoto: empty pic id")
	}

	pid := bson.ObjectIdHex(pidStr)
	picObj := getPhotoObj(pid, ctrl.photoCol)
	albumObj := getAlbumObj(picObj.Album, ctrl.albumCol)
	uid := ctrl.User.ObjectId

	if picObj.Owner == uid || albumObj.HostID == uid {
		removePhotoFile(pid.Hex(), uid.Hex())
		mgoRmFrmSet(ctrl.albumCol, picObj.Album, "photos", picObj.ObjectId)
		ifErr(ctrl.photoCol.RemoveId(pid))
	}

}
