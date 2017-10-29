package main

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestUserSearch(t *testing.T) {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	c := session.DB("test").C("accnts")

	names := []string{
		"Fred",
		"FredBurger",
		"Cat",
		"Dog",
	}

	joinedDates := "testing entry"

	for _, name := range names {
		ifErr(c.Insert(User{ObjectId: bson.NewObjectId(), Nickname: name, Joined: joinedDates}))
	}

	testLetter := 'f'
	fetchProfiles := GetProfilesMgo(string(testLetter))
	c.RemoveAll(bson.M{"joined": joinedDates})

	lenOfFetch := len(fetchProfiles)
	if lenOfFetch != 2 {
		t.Errorf("Matched with %d, should of been 2", lenOfFetch)
	}
	for _, profile := range fetchProfiles {
		nickname := profile.Nickname
		if []rune(nickname)[0] != testLetter {
			t.Errorf("Result was incorrect, got: %s, which does not start with %c", nickname, testLetter)
		}

	}

}
