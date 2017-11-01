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
	ctrl := getPDUController(r)
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
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.RemoveUser(r.PostFormValue("Password"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDMgoController()
	defer ctrl.session.Close()
	msg := new(LoginMsg)
	FillStruct(r, msg)
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
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ifErr(json.NewEncoder(w).
		Encode(
			GetFriendsResponse{ctrl.GetFriendReqs(), ctrl.GetFriendsMgo()}))
}

//SearchUser ...
func SearchUser(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	nameLike := pat.Param(r, "NAME_LIKE")
	fechedProfiles := ctrl.GetProfilesMgo(nameLike)
	ifErr(json.NewEncoder(w).Encode(fechedProfiles))
}

func AcceptReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.AcceptReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

// DeclineReq ...
func DeclineReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.DeclineReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

// SendReq ...
func SendReq(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.SendReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
