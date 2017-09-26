package main

import (
	//	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"log"
)

type User struct {
	bson.ObjectId               "_id"
	Username                    string
	Email                       string
	Password                    string
	Nickname                    string
	Joined                      string
	Friends, Albums, FriendReqs []bson.ObjectId
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

func InsertUser(msg AddUserMsg) bson.ObjectId {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("accnts")
	newUser := newUser(msg)

	ifErr(c.Insert(newUser))
	return newUser.ObjectId
}

func RemoveUser(msg DelUserMsg) {

	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

	bsonID := bson.ObjectIdHex(msg.Id)

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
		RemoveAlbum(AlbumMsgToken{AlbumId: album.Hex(), UserId: user.ObjectId.Hex()})
	}
	ifErr(c.RemoveId(bsonID))

}
