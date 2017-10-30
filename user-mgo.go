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

func (c *MgoController) checkUserExist(username, email string) CreateUserError {
	usernameFind := new(User)
	emailFind := new(User)

	c.userCol.Find(bson.M{"username": username}).One(usernameFind)
	c.userCol.Find(bson.M{"email": email}).One(emailFind)

	return CreateUserError{Username: (usernameFind.ObjectId.Hex() != ""), Email: (emailFind.ObjectId.Hex() != "")}
}

//InsertUser ...
func (c *MgoController) InsertUser(msg AddUserMsg) AddUserResp {
	newUser := newUser(msg)

	userDBCheck := c.checkUserExist(msg.Username, msg.Email)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(c.userCol.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

// RemoveUser ...
// TODO:	~Remove user from all guest lists
// 			~Remove user photos
func (c *MgoController) RemoveUser(msg DelUserMsg) {
	bsonID := bson.ObjectIdHex(msg.ID)

	var user User
	ifErr(c.userCol.FindId(bsonID).One(&user))

	if user.Password != msg.Password {
		panic("Wrong password")
	}

	for _, friend := range user.Friends {
		mgoRmFrmSetID(c.userCol, friend, "friends", user.ObjectId)
	}

	for _, album := range user.Albums {
		c.RemoveAlbum(AlbumMsgToken{AID: album.Hex(), UID: user.ObjectId.Hex()})
	}
	userPath := PrjDir + msg.ID
	ifErr(os.RemoveAll(userPath))
	ifErr(c.userCol.RemoveId(bsonID))
}

//GetUser ...
func (c *MgoController) GetUser(msg LoginMsg) *interface{} {
	foundUser := new(interface{})

	c.userCol.Find(bson.M{"username": msg.Username, "password": msg.Password}).One(foundUser)

	if foundUser == nil {
		panic("No user found")
	}
	return foundUser
}

//GetUserAlbums ...
func (c *MgoController) GetUserAlbums(uid string) GetAlbumsResp {
	foundUser := new(User)
	ifErr(c.userCol.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))

	list := foundUser.Albums
	taggedList := foundUser.Tagged
	created := make([]GetAlbumResp, len(list))
	album := new(Album)
	for i, albumID := range list {
		ifErr(c.albumCol.Find(bson.M{"_id": albumID}).One(album))
		created[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	tagged := make([]GetAlbumResp, len(taggedList))
	for i, albumID := range taggedList {
		ifErr(c.albumCol.Find(bson.M{"_id": albumID}).One(album))
		tagged[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	return GetAlbumsResp{Created: created, Tagged: tagged}
}

//GetFriendReqs ...
func (c *MgoController) GetFriendReqs(uid string) []SimpleUser {
	foundUser := new(User)
	ifErr(c.userCol.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))

	var results []SimpleUser
	foundFriendReq := new(User)
	for _, friendReqID := range foundUser.FriendReqs {
		ifErr(c.userCol.FindId(friendReqID).One(foundFriendReq))
		results = append(results, SimpleUser{Name: foundFriendReq.Nickname, ID: friendReqID.Hex()})
	}

	return results
}

//GetFriendsMgo ...
func (c *MgoController) GetFriendsMgo(uid string) []SimpleUser {
	foundUser := new(User)
	ifErr(c.userCol.FindId(bson.ObjectIdHex(uid)).One(foundUser))

	var results []SimpleUser
	foundFriend := new(User)
	for _, friendID := range foundUser.Friends {
		ifErr(c.userCol.FindId(friendID).One(foundFriend))
		results = append(results, SimpleUser{Name: foundFriend.Nickname, ID: friendID.Hex()})
	}

	return results
}

// GetProfilesMgo ...
func (c *MgoController) GetProfilesMgo(nameLike string) []UserProfile {
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
