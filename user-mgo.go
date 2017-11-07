package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

//InsertUser ...
func (ctrl *Controller) InsertUser(msg AddUserMsg) AddUserResp {
	newUser := ServerUser{
		ObjectId: bson.NewObjectId(),
		Username: msg.Username,
		Email:    msg.Email,
		Password: msg.Password,
		UserProfile: UserProfile{
			msg.Nickname,
			time.Now().UTC().String(),
		},
	}

	if msg.Username == "" || msg.Password == "" {
		panic("Unidentifiable user")
	}

	userDBCheck := checkIfUserExist(msg.Username, msg.Email, ctrl.userCol)
	// fmt.Println("Checks from db~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// fmt.Println(userDBCheck)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(ctrl.userCol.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

// RemoveUser ...
// TODO:	~Remove user from all guest lists
// 			~Remove user photos
func (ctrl *Controller) RemoveUser(password string) {
	user := ctrl.ServerUser

	if user.Password != password {
		panic(fmt.Sprintf("Can not remove user with %s", password))
	}

	for _, friend := range user.Friends {
		mgoRmFrmSet(ctrl.userCol, friend, "friends", user.ObjectId)
	}

	for _, album := range user.Albums {
		ctrl.removeAlbumInternal(album)
	}
	userPath := PrjDir + user.ObjectId.Hex()
	ifErr(os.RemoveAll(userPath))
	ifErr(ctrl.userCol.RemoveId(user.ObjectId))
}

//GetUser ...
func (ctrl *Controller) GetUser(username, password string) ClientUser {
	var foundUser ServerUser
	ifErr(ctrl.userCol.Find(bson.M{"username": username, "password": password}).One(&foundUser))

	if foundUser.ObjectId.Hex() == "" {
		fmt.Println(fmt.Sprintf("No user found with username: %s and password: %s", username, password))
	}

	return ClientUser{
		ObjectID:           foundUser.ObjectId,
		Email:              foundUser.Email,
		UserProfile:        foundUser.UserProfile,
		GetAlbumsResp:      getAlbumsRespInternal(foundUser, ctrl.albumCol),
		GetFriendsResponse: getFriendsRespInternal(foundUser, ctrl.userCol),
	}
}

func getFriendsRespInternal(user ServerUser, userCol *mgo.Collection) GetFriendsResponse {
	return GetFriendsResponse{
		FriendReqs: getUserNicknamesInternal(user.FriendReqs, userCol),
		Friends:    getUserNicknamesInternal(user.FriendReqs, userCol),
	}
}

//GetFriendReqs ...
func (ctrl *Controller) GetFriendReqs() []string {
	return getUserNicknamesInternal(ctrl.ServerUser.FriendReqs, ctrl.albumCol)
}

func getUserNicknamesInternal(uids []bson.ObjectId, userCol *mgo.Collection) []string {
	var results []string
	for _, friendReqID := range uids {
		results = append(results, getUserObj(friendReqID, userCol).Nickname)
	}

	return results

}

//GetFriendsMgo ...
func (ctrl *Controller) GetFriendsMgo() []string {
	return getUserNicknamesInternal(ctrl.Friends, ctrl.userCol)

}

// GetProfilesMgo ...
func (ctrl *Controller) GetProfilesMgo(nameLike string) []UserProfile {
	var real []UserProfile

	ifErr(ctrl.userCol.EnsureIndexKey("nickname"))
	ifErr(
		ctrl.userCol.Find(
			bson.M{"nickname": &bson.RegEx{Pattern: "^" + nameLike, Options: "i"}}).
			Select(
				bson.M{"joined": 1, "nickname": 1}).
			All(&real))

	return real
}

func checkIfUserExist(username, email string, userCol *mgo.Collection) CreateUserError {
	usernameFind, err := userCol.Find(bson.M{"username": username}).Count()
	ifErr(err)
	emailFind, err := userCol.Find(bson.M{"email": email}).Count()
	ifErr(err)

	return CreateUserError{Username: usernameFind > 0, Email: emailFind > 0}
}
