package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func newUser(msg AddUserMsg) User {
	newU := User{Joined: time.Now().UTC().String()}
	newU.ObjectId = bson.NewObjectId()
	newU.Nickname = msg.Nickname
	newU.Email = msg.Email
	newU.Username = msg.Username
	newU.Password = msg.Password
	return newU
}

//InsertUser ...
func (c *Controller) InsertUser(msg AddUserMsg) AddUserResp {
	newUser := newUser(msg)

	userDBCheck := checkIfUserExist(msg.Username, msg.Email, c.userCol)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(c.userCol.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

// RemoveUser ...
// TODO:	~Remove user from all guest lists
// 			~Remove user photos
func (ctrl *Controller) RemoveUser(password string) {
	user := ctrl.User

	if user.Password != password {
		panic(fmt.Sprintf("Can not remove user with %s", password))
	}

	for _, friend := range user.Friends {
		mgoRmFrmSet(ctrl.userCol, friend, "friends", user.ObjectId)
	}

	for _, album := range user.Albums {
		ctrl.RemoveAlbumInternal(album)
	}
	userPath := PrjDir + user.ObjectId.Hex()
	ifErr(os.RemoveAll(userPath))
	ifErr(ctrl.userCol.RemoveId(user.ObjectId))
}

//GetUser ...
func (ctrl *Controller) GetUser(username, password string) interface{} {
	var foundUser interface{}
	ctrl.userCol.Find(bson.M{"username": username, "password": password}).One(&foundUser)

	if foundUser == nil {
		panic(fmt.Sprintf("No user found with username: %s and password: %s", username, password))
	}
	return foundUser
}

//GetFriendReqs ...
func (ctrl *Controller) GetFriendReqs() []string {

	var results []string
	for _, friendReqID := range ctrl.User.FriendReqs {
		results = append(results, getUserObj(friendReqID, ctrl.userCol).Nickname)
	}

	return results
}

//GetFriendsMgo ...
func (ctrl *Controller) GetFriendsMgo() []string {

	var results []string
	for _, friendID := range ctrl.User.Friends {
		results = append(results, getUserObj(friendID, ctrl.userCol).Nickname)
	}

	return results
}

// GetProfilesMgo ...
func (c *Controller) GetProfilesMgo(nameLike string) []UserProfile {
	var real []UserProfile

	ifErr(c.userCol.EnsureIndexKey("nickname"))
	ifErr(
		c.userCol.Find(
			bson.M{"nickname": &bson.RegEx{Pattern: "^" + nameLike, Options: "i"}}).
			Select(
				bson.M{"joined": 1, "nickname": 1}).
			All(&real))

	return real
}

// AcceptReqMgo ...
func (ctrl *Controller) AcceptReqMgo(fuidStr string) {
	fuid := bson.ObjectIdHex(fuidStr)
	uid := ctrl.User.ObjectId
	mgoRmFrmSet(ctrl.userCol, uid, "friendReqs", fuid)
	mgoAddToSet(ctrl.userCol, uid, "friends", fuid)
	mgoAddToSet(ctrl.userCol, fuid, "friends", uid)
}

// DeclineReqRequest ...
func (ctrl *Controller) DeclineReqMgo(fuidStr string) {
	mgoRmFrmSet(ctrl.userCol, ctrl.User.ObjectId, "friendReqs", bson.ObjectIdHex(fuidStr))
}

// SendReq ...
func (ctrl *Controller) SendReqMgo(fuidStr string) {
	mgoAddToSet(ctrl.userCol, bson.ObjectIdHex(fuidStr), "friendReqs", ctrl.User.ObjectId)
}

// RemoveFriend ...
func (ctrl *Controller) RemoveFriendMgo(fuidStr string) {
	uid := ctrl.User.ObjectId
	fuid := bson.ObjectIdHex(fuidStr)
	mgoRmFrmSet(ctrl.userCol, uid, "friends", fuid)
	mgoRmFrmSet(ctrl.userCol, fuid, "friends", uid)
}

func checkIfUserExist(username, email string, userCol *mgo.Collection) CreateUserError {
	usernameFind, err := userCol.Find(bson.M{"username": username}).Count()
	ifErr(err)
	emailFind, err := userCol.Find(bson.M{"email": email}).Count()
	ifErr(err)

	return CreateUserError{Username: usernameFind > 0, Email: emailFind > 0}
}
