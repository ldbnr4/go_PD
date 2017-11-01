package main

import (
	"os"
	"time"

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

func (c *PDMgoController) checkIfUserExist(username, email string) CreateUserError {
	usernameFind := new(User)
	emailFind := new(User)

	c.userCol.Find(bson.M{"username": username}).One(usernameFind)
	c.userCol.Find(bson.M{"email": email}).One(emailFind)

	return CreateUserError{Username: (usernameFind.ObjectId.Hex() != ""), Email: (emailFind.ObjectId.Hex() != "")}
}

//InsertUser ...
func (c *PDMgoController) InsertUser(msg AddUserMsg) AddUserResp {
	newUser := newUser(msg)

	userDBCheck := c.checkIfUserExist(msg.Username, msg.Email)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(c.userCol.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

// RemoveUser ...
// TODO:	~Remove user from all guest lists
// 			~Remove user photos
func (pdMgo *PDUIDMgoController) RemoveUser(password string) {
	var user User
	ifErr(pdMgo.userCol.FindId(pdMgo.UID).One(&user))

	if user.Password != msg.Password {
		panic("Wrong password")
	}

	for _, friend := range user.Friends {
		mgoRmFrmSet(pdMgo.userCol, friend, "friends", user.ObjectId)
	}

	for _, album := range user.Albums {
		pdMgo.RemoveAlbum(AlbumMsgToken{AID: album.Hex()})
	}
	userPath := PrjDir + pdMgo.UID.Hex()
	ifErr(os.RemoveAll(userPath))
	ifErr(c.userCol.RemoveId(c.UID))
}

//GetUser ...
func (c *PDMgoController) GetUser(msg LoginMsg) *interface{} {
	foundUser := new(interface{})

	c.userCol.Find(bson.M{"username": msg.Username, "password": msg.Password}).One(foundUser)

	if foundUser == nil {
		panic("No user found")
	}
	return foundUser
}

//GetFriendReqs ...
func (c *PDUIDMgoController) GetFriendReqs(uid bson.ObjectId) []string {
	foundUser := new(User)
	ifErr(c.userCol.Find(bson.M{"_id": uid}).One(foundUser))

	var results []string
	foundFriendReq := new(User)
	for _, friendReqID := range foundUser.FriendReqs {
		ifErr(c.userCol.FindId(friendReqID).One(foundFriendReq))
		results = append(results, foundFriendReq.Nickname)
	}

	return results
}

//GetFriendsMgo ...
func (c *PDUIDMgoController) GetFriendsMgo(uid string) []string {
	foundUser := new(User)
	ifErr(c.userCol.FindId(bson.ObjectIdHex(uid)).One(foundUser))

	var results []string
	foundFriend := new(User)
	for _, friendID := range foundUser.Friends {
		ifErr(c.userCol.FindId(friendID).One(foundFriend))
		results = append(results, foundFriend.Nickname)
	}

	return results
}

// GetProfilesMgo ...
func (c *PDUIDMgoController) GetProfilesMgo(nameLike string) []UserProfile {
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
func (c *PDUIDMgoController) AcceptReqMgo(msg FriendReqRequest) {
	mgoRmFrmSet(c.userCol, c.UID, "friendReqs", bson.ObjectIdHex(msg.FriendUID))
	mgoAddToSet(c.userCol, c.UID, "friends", bson.ObjectIdHex(msg.FriendUID))
	mgoAddToSet(c.userCol, bson.ObjectIdHex(msg.FriendUID), "friends", c.UID)
}

// DeclineReqRequest ...
func (c *PDUIDMgoController) DeclineReqMgo(msg FriendReqRequest) {
	mgoRmFrmSet(c.userCol, c.UID, "friendReqs", bson.ObjectIdHex(msg.FriendUID))
}

// SendReq ...
func (c *PDUIDMgoController) SendReqMgo(msg FriendReqRequest) {
	mgoAddToSet(c.userCol, bson.ObjectIdHex(msg.FriendUID), "friendReqs", c.UID)
}

// RemoveFriend ...
func (c *PDUIDMgoController) RemoveFriendMgo(msg FriendReqRequest) {
	mgoRmFrmSet(c.userCol, c.UID, "friends", bson.ObjectIdHex(msg.FriendUID))
	mgoRmFrmSet(c.userCol, bson.ObjectIdHex(msg.FriendUID), "friends", c.UID)
}
