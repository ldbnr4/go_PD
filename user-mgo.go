package main

import (
	//	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	//	"log"
)

type User struct {
	bson.ObjectId `_id`
	AddUserMsg
	Joined                      string
	Friends, Albums, FriendReqs []string
}

func newUser(msg AddUserMsg) User {
	newU := User{AddUserMsg: msg, Joined: time.Now().UTC().String()}
	newU.ObjectId = bson.NewObjectId()
	return newU
}

func InsertUser(msg AddUserMsg) bson.ObjectId {
	session, err := mgo.Dial("localhost:27012")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("accnts")
	newUser := newUser(msg)

	err = c.Insert(newUser)
	if err != nil {
		panic(err)
	}

	return newUser.ObjectId
}
