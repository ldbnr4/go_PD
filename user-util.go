package main

import (
	"io"
	"os"
)

func setUpUserDirectory(uid string) {
	userDirLocation := PrjDir + uid
	ifErr(os.MkdirAll(userDirLocation, os.ModePerm))
	setDefaultProfilePic(uid)
}

// setDefaultProfilePic copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func setDefaultProfilePic(uid string) {
	src := PrjDir + "profile.png"
	dst := PrjDir + uid + "/" + uid

	in, err := os.Open(src)
	ifErr(err)

	out, err := os.Create(dst)
	ifErr(err)
	// TODO: Error check that compares bytes written to original btyes of file
	_, err = io.Copy(out, in)
	ifErr(err)
	ifErr(out.Sync())
	defer ifErr(in.Close())
	defer ifErr(out.Close())
}

func (ctrl *Controller) convertServerToClient(user ServerUser) ClientUser {
	return ClientUser{
		ObjectID:           user.ObjectId,
		Email:              user.Email,
		UserProfile:        user.UserProfile,
		GetAlbumsResp:      getAlbumsRespInternal(user, ctrl.albumCol),
		GetFriendsResponse: getFriendsRespInternal(user, ctrl.userCol),
	}
}
