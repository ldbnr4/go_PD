package main

import "gopkg.in/mgo.v2/bson"

func onTheGuestList(album Album, visitor bson.ObjectId) bool {
	allowed := false
	for _, guest := range album.Guest {
		if guest == visitor {
			allowed = true
			break
		}
	}
	return allowed
}
