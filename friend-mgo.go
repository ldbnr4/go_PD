package main

import "gopkg.in/mgo.v2/bson"

func (ctrl *Controller) AcceptReqMgo(fuidStr string) {
	fuid := bson.ObjectIdHex(fuidStr)
	uid := ctrl.User.ObjectId
	mgoRmFrmSet(ctrl.userCol, uid, "friendReqs", fuid)
	mgoAddToSet(ctrl.userCol, uid, "friends", fuid)
	mgoAddToSet(ctrl.userCol, fuid, "friends", uid)
}

func (ctrl *Controller) DeclineReqMgo(fuidStr string) {
	mgoRmFrmSet(ctrl.userCol, ctrl.User.ObjectId, "friendReqs", bson.ObjectIdHex(fuidStr))
}

func (ctrl *Controller) SendReqMgo(fuidStr string) {
	mgoAddToSet(ctrl.userCol, bson.ObjectIdHex(fuidStr), "friendReqs", ctrl.User.ObjectId)
}

func (ctrl *Controller) RemoveFriendMgo(fuidStr string) {
	uid := ctrl.User.ObjectId
	fuid := bson.ObjectIdHex(fuidStr)
	mgoRmFrmSet(ctrl.userCol, uid, "friends", fuid)
	mgoRmFrmSet(ctrl.userCol, fuid, "friends", uid)
}
