package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO: make first parameter type{mgo.Collection, bson.ObjectId}
//{ $addToSet: { <field1>: <value1>, (TODO:)... } }
func mgoAddToSet(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$addToSet"] = op

	ifErr(collection.UpdateId(id, update))

}

// TODO: make first parameter type{mgo.Collection, bson.ObjectId}
//{ $pull: { <field1>: <value|(TODO:)condition>} }
func mgoRmFrmSet(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$pull"] = op

	ifErr(collection.UpdateId(id, update))

}

func convertToObjectID(idStr string) bson.ObjectId {
	return bson.ObjectIdHex(idStr)
}
