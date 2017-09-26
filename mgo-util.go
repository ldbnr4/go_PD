package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//{ $addToSet: { <field1>: <value1>, (TODO:)... } }
func mgoAddToSetID(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$addToSet"] = op

	ifErr(collection.UpdateId(id, update))

}

//{ $pull: { <field1>: <value|(TODO:)condition>} }
func mgoRmFrmSetID(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$pull"] = op

	ifErr(collection.UpdateId(id, update))

}
