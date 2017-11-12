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
	return AddUserResp{ClientUser: ctrl.convertServerToClient(newUser), Error: userDBCheck}
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

	return ctrl.convertServerToClient(foundUser)
}

func getFriendsRespInternal(user ServerUser, userCol *mgo.Collection) GetFriendsResponse {
	return GetFriendsResponse{
		FriendReqs: getUserProfiles(user.FriendReqs, userCol),
		Friends:    getUserProfiles(user.FriendReqs, userCol),
	}
}

//GetFriendReqs ...
func (ctrl *Controller) GetFriendReqs() []UserProfile {
	return getUserProfiles(ctrl.ServerUser.FriendReqs, ctrl.userCol)
}

func getUserProfiles(uids []bson.ObjectId, userCol *mgo.Collection) []UserProfile {
	var results []UserProfile
	for _, friendReqID := range uids {
		results = append(results, getUserObj(friendReqID, userCol).UserProfile)
	}
	return results
}

//GetFriendsMgo ...
func (ctrl *Controller) GetFriendsMgo() []UserProfile {
	return getUserProfiles(ctrl.Friends, ctrl.userCol)

}

// GetProfilesMgo ...
func (ctrl *Controller) GetProfilesMgo(nameLike string) []UserProfile {
	var buffer []map[string]map[string]string

	ifErr(ctrl.userCol.EnsureIndexKey("userprofile.nickname"))
	ifErr(
		ctrl.userCol.Find(
			bson.M{"userprofile.nickname": bson.RegEx{Pattern: "^" + nameLike, Options: "i"}}).
			Select(
				bson.M{"userprofile.joined": true, "userprofile.nickname": true}).
			All(&buffer))
	var real []UserProfile
	for _, profile := range buffer {
		real = append(real, UserProfile{
			profile["userprofile"]["nickname"],
			profile["userprofile"]["joined"],
		})
	}

	return real
}

func checkIfUserExist(username, email string, userCol *mgo.Collection) CreateUserError {
	usernameFind, err := userCol.Find(bson.M{"username": username}).Count()
	ifErr(err)
	emailFind, err := userCol.Find(bson.M{"email": email}).Count()
	ifErr(err)

	return CreateUserError{Username: usernameFind > 0, Email: emailFind > 0}
}

func (ctrl Controller) getUIDFromNickname(nickname string) bson.ObjectId {
	uid := make(map[string]bson.ObjectId)
	ifErr(ctrl.userCol.Find(bson.M{"userprofile.nickname": nickname}).Select(bson.M{"_id": 1}).One(&uid))
	return uid["_id"]
}
