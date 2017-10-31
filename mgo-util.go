package main

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO: Use this type to handle all database interaction
type MgoController struct {
	session                     *mgo.Session
	userCol, albumCol, photoCol *mgo.Collection
	UID                         bson.ObjectId
}

// func (c *Controller) getUser(c web.C, w http.ResponseWriter, r *http.Request) {
//     // Use the session var here
// }
// func main() {
//     ctl, err := NewController()
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer ctl.session.Close()
//     goji.Get("/user", ctl.getUser)
//     goji.Serve()
// }
func getController(r *http.Request) *MgoController {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	db := session.DB("test")
	retCtrl := &MgoController{
		session:  session,
		userCol:  db.C("accnts"),
		albumCol: db.C("albums"),
		photoCol: db.C("photos"),
	}
	uidStr := r.Header.Get("UID")
	if uidStr != "" {
		retCtrl.UID = bson.ObjectIdHex(uidStr)
		count, err := retCtrl.userCol.FindId(retCtrl.UID).Count()
		ifErr(err)
		if count != 1 {
			panic("Invalid UID")
		}
	}
	return retCtrl
}

// TODO: make first parameter type{mgo.Collection, bson.ObjectId}
//{ $addToSet: { <field1>: <value1>, (TODO:)... } }
func mgoAddToSetID(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$addToSet"] = op

	ifErr(collection.UpdateId(id, update))

}

// TODO: make first parameter type{mgo.Collection, bson.ObjectId}
//{ $pull: { <field1>: <value|(TODO:)condition>} }
func mgoRmFrmSetID(collection *mgo.Collection, id bson.ObjectId, field string, value interface{}) {

	update := make(map[string]map[string]interface{})
	op := make(map[string]interface{})
	op[field] = value
	update["$pull"] = op

	ifErr(collection.UpdateId(id, update))

}
