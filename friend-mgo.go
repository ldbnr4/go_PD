package main

func (ctrl *Controller) AcceptReqMgo(nicknameStr string) {
	fuid := ctrl.getUIDFromNickname(nicknameStr)
	uid := ctrl.ServerUser.ObjectId
	mgoRmFrmSet(ctrl.userCol, uid, "friendreqs", fuid)
	mgoAddToSet(ctrl.userCol, uid, "friends", fuid)
	mgoAddToSet(ctrl.userCol, fuid, "friends", uid)
}

func (ctrl *Controller) DeclineReqMgo(nicknameStr string) {
	fuid := ctrl.getUIDFromNickname(nicknameStr)
	mgoRmFrmSet(ctrl.userCol, ctrl.ServerUser.ObjectId, "friendreqs", fuid)
}

func (ctrl *Controller) SendReqMgo(nicknameStr string) {
	fuid := ctrl.getUIDFromNickname(nicknameStr)
	mgoAddToSet(ctrl.userCol, fuid, "friendreqs", ctrl.ServerUser.ObjectId)
}

func (ctrl *Controller) RemoveFriendMgo(nicknameStr string) {
	uid := ctrl.ServerUser.ObjectId
	fuid := ctrl.getUIDFromNickname(nicknameStr)
	mgoRmFrmSet(ctrl.userCol, uid, "friends", fuid)
	mgoRmFrmSet(ctrl.userCol, fuid, "friends", uid)
}

func (ctrl *Controller) getNetworkRelation(nickname string) string {
	uid := ctrl.getUIDFromNickname(nickname)
	// panic(uid)
	freindUser := getUserObj(uid, ctrl.userCol)

	if uid == ctrl.ObjectId {
		return "OWN"
	}

	for _, friend := range freindUser.Friends {
		if friend == ctrl.ObjectId {
			return "FRND"
		}
	}

	for _, fiendReq := range freindUser.FriendReqs {
		if fiendReq == ctrl.ObjectId {
			return "PEND"
		}
	}

	for _, friendReq := range ctrl.FriendReqs {
		if friendReq == freindUser.ObjectId {
			return "PEND_ACTION"
		}
	}

	return "ISO"

}
