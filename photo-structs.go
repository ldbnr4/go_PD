package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Photo struct {
	bson.ObjectId "_id"
	Upload        time.Time
	Owner         bson.ObjectId
	Album         bson.ObjectId
}

type PhotoIDMsg struct {
	PID string
}
