package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

//InsertUser ...
func (c *Controller) InsertUser(msg AddUserMsg) AddUserResp {
	newUser := User{
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

	userDBCheck := checkIfUserExist(msg.Username, msg.Email, c.userCol)
	// fmt.Println("Checks from db~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// fmt.Println(userDBCheck)

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
		ctrl.removeAlbumInternal(album)
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
		fmt.Sprintf("No user found with username: %s and password: %s", username, password)
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

func checkIfUserExist(username, email string, userCol *mgo.Collection) CreateUserError {
	usernameFind, err := userCol.Find(bson.M{"username": username}).Count()
	ifErr(err)
	emailFind, err := userCol.Find(bson.M{"email": email}).Count()
	ifErr(err)

	return CreateUserError{Username: usernameFind > 0, Email: emailFind > 0}
}
