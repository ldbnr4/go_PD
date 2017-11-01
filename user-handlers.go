package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"goji.io/pat"
)

//UserCreate ...
func UserCreate(w http.ResponseWriter, r *http.Request) {
	msg := new(AddUserMsg)
	FillStruct(r, msg)

	ctrl := getPDController(r)
	defer ctrl.session.Close()

	insertResp := ctrl.InsertUser(*msg)
	setUpUserDirectory(insertResp.ID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(insertResp))
}

//UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelUserMsg)
	FillStruct(r, msg)
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	ctrl.RemoveUser(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	msg := new(LoginMsg)
	FillStruct(r, msg)
	ctrl := getExposedController(r)
	defer ctrl.session.Close()
	ifErr(json.NewEncoder(w).Encode(ctrl.GetUser(*msg)))
}

//ProfPic ...
func ProfPic(w http.ResponseWriter, r *http.Request) {
	UID := pat.Param(r, "UID")

	// TODO: check if UID is a valid UID
	f, err := os.Open(PrjDir + UID + "/" + UID)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}

//GetFriends ...
func GetFriends(w http.ResponseWriter, r *http.Request) {
	UID := pat.Param(r, "UID")
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	ifErr(json.NewEncoder(w).
		Encode(
			GetFriendsResponse{ctrl.GetFriendReqs(UID), ctrl.GetFriendsMgo(UID)}))
}

//GetAlbums ...
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	UID := pat.Param(r, "UID")
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	aids := ctrl.GetUserAlbums(UID)
	ifErr(json.NewEncoder(w).Encode(aids))
}

//SearchUser ...
func SearchUser(w http.ResponseWriter, r *http.Request) {
	nameLike := pat.Param(r, "NAME_LIKE")
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	fechedProfiles := ctrl.GetProfilesMgo(nameLike)
	ifErr(json.NewEncoder(w).Encode(fechedProfiles))
}

func AcceptReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	var msg FriendReqRequest
	FillStruct(r, msg)
	ctrl.AcceptReqMgo(msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

// DeclineReq ...
func DeclineReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	var msg FriendReqRequest
	FillStruct(r, msg)
	ctrl.DeclineReqMgo(msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

// SendReq ...
func SendReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDController(r)
	defer ctrl.session.Close()
	var msg FriendReqRequest
	FillStruct(r, msg)
	ctrl.SendReqMgo(msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
