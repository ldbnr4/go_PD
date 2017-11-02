package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MgoCollections struct {
	userCol, albumCol, photoCol *mgo.Collection
}

type Controller struct {
	session *mgo.Session
	MgoCollections
	User
}

func getController(r *http.Request) Controller {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	envStr := r.Header.Get("ENV")
	db := session.DB(envStr)
	mgoCols := MgoCollections{
		userCol:  db.C("accnts"),
		albumCol: db.C("albums"),
		photoCol: db.C("photos"),
	}
	ctrl := Controller{
		session:        session,
		MgoCollections: mgoCols,
	}

	if !isExposedPath(r.URL.Path) {
		uidStr := r.Header.Get("UID")
		if uidStr == "" {
			panic("This route requires identification!")
		}
		ctrl.User = getUserObj(bson.ObjectIdHex(uidStr), mgoCols.userCol)
	}
	return ctrl
}

func isExposedPath(path string) bool {
	exposedPaths := []string{
		"/login",
		"/user",
	}
	for _, ePath := range exposedPaths {
		if ePath == path {
			return true
		}
	}
	return false
}
