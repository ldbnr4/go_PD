package main

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO: Use this type to handle all database interaction
// =============================================================================================================================================================
// type Controller struct {
//     session *Session
// }

// func NewController() (*Controller, error) {
//     if uri := os.Getenv("MONGOHQ_URL"); uri == "" {
//         return nil, fmt.Errorf("no DB connection string provided")
//     }
//     session, err := mgo.Dial(uri)
//     if err != nil {
//         return nil, err
//     }
//     return &Controller{
//         session: session,
//     }, nil
// }

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
// =============================================================================================================================================================
// type DataStore struct {
//     session *mgo.Session
// }

// func (ds *DataStore) ucol() *mgo.Collection { ... }

// func (ds *DataStore) UserExist(user string) bool { ... }
// There are many benefits to that design.
// An important one is that it allows you to have multiple sessions in flight at the same time,
// so if you have an http handler, for example, you can create a local session that is backed by an independent session just for that one request:

// func (s *WebSite) dataStore() *DataStore {
//     return &DataStore{s.session.Copy()}
// }

// func (s *WebSite) HandleRequest(...) {
//     ds := s.dataStore()
//     defer ds.Close()
//     ...
// }

func newUser(msg AddUserMsg) User {
	newU := User{Joined: time.Now().UTC().String()}
	newU.ObjectId = bson.NewObjectId()
	newU.Nickname = msg.Nickname
	newU.Email = msg.Email
	newU.Username = msg.Username
	newU.Password = msg.Password
	return newU
}

func checkUserExist(username string, email string) CreateUserError {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")

	usernameFind := new(User)
	emailFind := new(User)

	c.Find(bson.M{"username": username}).One(usernameFind)
	c.Find(bson.M{"email": email}).One(emailFind)

	return CreateUserError{Username: (usernameFind.ObjectId.Hex() != ""), Email: (emailFind.ObjectId.Hex() != "")}
}

//InsertUser ...
func InsertUser(msg AddUserMsg) AddUserResp {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")
	newUser := newUser(msg)

	userDBCheck := checkUserExist(msg.Username, msg.Email)

	if !userDBCheck.Username && !userDBCheck.Email {
		ifErr(c.Insert(newUser))
	}
	return AddUserResp{ID: newUser.ObjectId.Hex(), Error: userDBCheck}
}

func RemoveUser(msg DelUserMsg) {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

	bsonID := bson.ObjectIdHex(msg.ID)

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
		RemoveAlbum(AlbumMsgToken{AID: album.Hex(), UID: user.ObjectId.Hex()})
	}
	userPath := PrjDir + msg.ID
	ifErr(os.RemoveAll(userPath))
	ifErr(c.RemoveId(bsonID))
}

//RetrieveUser ...
func RetrieveUser(msg LoginMsg) *interface{} {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()

	c := session.DB("test").C("accnts")

	foundUser := new(interface{})

	c.Find(bson.M{"username": msg.Username, "password": msg.Password}).One(foundUser)
	// panic(*foundUser)

	if foundUser == nil {
		panic("No user found")
	}
	return foundUser
}

//GetUserAlbums ...
func GetUserAlbums(uid string) GetAlbumsResp {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	c := session.DB("test").C("accnts")
	foundUser := new(User)
	ifErr(c.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))

	a := session.DB("test").C("albums")

	list := foundUser.Albums
	taggedList := foundUser.Tagged
	created := make([]GetAlbumResp, len(list))
	album := new(Album)
	for i, albumID := range list {
		ifErr(a.Find(bson.M{"_id": albumID}).One(album))
		created[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	tagged := make([]GetAlbumResp, len(taggedList))
	for i, albumID := range taggedList {
		ifErr(a.Find(bson.M{"_id": albumID}).One(album))
		tagged[i] = GetAlbumResp{Title: album.Title, ID: albumID.Hex()}
	}

	return GetAlbumsResp{Created: created, Tagged: tagged}
}

//GetFriendReqs ...
func GetFriendReqs(uid string) []SimpleUser {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	c := session.DB("test").C("accnts")
	foundUser := new(User)
	ifErr(c.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))

	var results []SimpleUser
	foundFriendReq := new(User)
	for _, friendReqID := range foundUser.FriendReqs {
		ifErr(c.Find(bson.M{"_id": friendReqID}).One(foundFriendReq))
		results = append(results, SimpleUser{Name: foundFriendReq.Nickname, ID: friendReqID.Hex()})
	}

	return results
}

//GetFriendsMgo ...
func GetFriendsMgo(uid string) []SimpleUser {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	c := session.DB("test").C("accnts")
	foundUser := new(User)
	ifErr(c.Find(bson.M{"_id": bson.ObjectIdHex(uid)}).One(foundUser))

	var results []SimpleUser
	foundFriend := new(User)
	for _, friendID := range foundUser.Friends {
		ifErr(c.Find(bson.M{"_id": friendID}).One(foundFriend))
		results = append(results, SimpleUser{Name: foundFriend.Nickname, ID: friendID.Hex()})
	}

	return results
}

// GetProfilesMgo ...
func GetProfilesMgo(nameLike string) []UserProfile {
	session, err := mgo.Dial("localhost:27012")
	ifErr(err)
	defer session.Close()
	c := session.DB("test").C("accnts")
	var real []UserProfile

	ifErr(c.EnsureIndexKey("nickname"))
	ifErr(
		c.Find(
			bson.M{"nickname": &bson.RegEx{Pattern: "^" + nameLike, Options: "i"}}).
			Select(
				bson.M{"joined": 1, "nickname": 1}).
			All(&real))

	return real
}
