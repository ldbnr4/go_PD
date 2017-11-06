package main

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

//UserCreate ...
func UserCreate(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	msg := new(AddUserMsg)
	FillStruct(r, msg)

	insertResp := ctrl.InsertUser(*msg)
	setUpUserDirectory(insertResp.ID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(insertResp))
}

//UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.RemoveUser(r.PostFormValue("Password"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ifErr(json.NewEncoder(w).Encode(ctrl.GetUser(
		r.FormValue("username"),
		r.FormValue("password"),
	)))
}

//ProfPic ...
func ProfPic(w http.ResponseWriter, r *http.Request) {
	UID := pat.Param(r, "UID")

	profPicPath := PrjDir + UID + "/" + UID
	serveFile(profPicPath, w)
}

//GetFriends ...
func GetFriends(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ifErr(json.NewEncoder(w).
		Encode(
			GetFriendsResponse{ctrl.GetFriendReqs(), ctrl.GetFriendsMgo()}))
}

//SearchUser ...
func SearchUser(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	nameLike := pat.Param(r, "NAME_LIKE")
	fechedProfiles := ctrl.GetProfilesMgo(nameLike)
	ifErr(json.NewEncoder(w).Encode(fechedProfiles))
}
