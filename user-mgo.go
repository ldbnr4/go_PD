package main

import (
	//	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"log"
)

//User ...
type User struct {
	bson.ObjectId                       "_id"
	Username                            string
	Email                               string
	Password                            string
	Nickname                            string
	Joined                              string
	Friends, Albums, FriendReqs, Tagged []bson.ObjectId
}

//CreateUserError ...
type CreateUserError struct {
	Username bool
	Email    bool
}

func newUser(msg AddUserMsg) User {
	newU := User{Joined: time.Now().UTC().String()}
	newU.ObjectId = bson.NewObjectId()
	newU.Nickname = msg.Nickname
	newU.Email = msg.Email
	newU.Username = msg.Username
	newU.Password = msg.Password
	return newU
}

func checkUserExist(username string, email string) CreateUserError {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")

	usernameFind := new(User)
	emailFind := new(User)

	c.Find(bson.M{"username": username}).One(usernameFind)
	c.Find(bson.M{"email": email}).One(emailFind)

	return CreateUserError{Username: (usernameFind.ObjectId.Hex() != ""), Email: (emailFind.ObjectId.Hex() != "")}
}

//InsertUser ...
func InsertUser(msg AddUserMsg) AddUserResp {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")
	newUser := newUser(msg)

	userDBCheck := checkUserExist(msg.Username, msg.Email)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(c.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

func RemoveUser(msg DelUserMsg) {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

	bsonID := bson.ObjectIdHex(msg.ID)

	var user User

	c := session.DB("test").C("accnts")
	ifErr(c.FindId(bsonID).One(&user))

	if user.Password != msg.Password {
		panic("Wrong password")
	}

	for _, friend := range user.Friends {
		mgoRmFrmSetID(c, friend, "friends", user.ObjectId)
		// element is the element from someSlice for where we are
	}

	for _, album := range user.Albums {
		RemoveAlbum(AlbumMsgToken{AID: album.Hex(), UID: user.ObjectId.Hex()})
	}
	userPath := PrjDir + msg.ID
	ifErr(os.RemoveAll(userPath))
	ifErr(c.RemoveId(bsonID))
}

//RetrieveUser ...
func RetrieveUser(msg LoginMsg) *interface{} {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")

	foundUser := new(interface{})

	c.Find(bson.M{"username": msg.Username, "password": msg.Password}).One(foundUser)
	// panic(*foundUser)

	if foundUser == nil {
		panic("No user found")
	}
	return foundUser
}

//GetUserAlbums ...
func GetUserAlbums(uid string) GetAlbumsResp {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")

	foundUser := new(User)

	ifErr(c.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))
	a := session.DB("test").C("albums")

	list := foundUser.Albums
	taggedList := foundUser.Tagged
	created := make([]GetAlbumResp, len(list))
	album := new(Album)
	for i, albumID := range list {
		ifErr(a.Find(bson.M{"_id": albumID}).One(album))
		created[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	tagged := make([]GetAlbumResp, len(taggedList))
	for i, albumID := range taggedList {
		ifErr(a.Find(bson.M{"_id": albumID}).One(album))
		tagged[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	return GetAlbumsResp{Created: created, Tagged: tagged}
}
