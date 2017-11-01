package main

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PDMgoController struct {
	session                     *mgo.Session
	userCol, albumCol, photoCol *mgo.Collection
}

// TODO: Use this type to handle all database interaction
type PDUMgoController struct {
	PDMgoController
	User
}

type ControllerParam struct {
	r        *http.Request
	identify bool
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
func getPDMgoController() PDMgoController {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	db := session.DB("test")
	return PDMgoController{
		session:  session,
		userCol:  db.C("accnts"),
		albumCol: db.C("albums"),
		photoCol: db.C("photos"),
	}
}

func getPDUController(r *http.Request) PDUMgoController {
	mgoCtrl := getPDMgoController()
	uidStr := r.Header.Get("UID")
	if uidStr == "" {
		panic("This route requires identification")
	}
	UID := bson.ObjectIdHex(uidStr)
	var user User
	ifErr(mgoCtrl.userCol.FindId(UID).One(user))
	return PDUIDMgoController{mgoCtrl, user}
}

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
