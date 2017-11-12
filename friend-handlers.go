package main

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func AcceptFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.AcceptReqMgo(r.PostFormValue("FRIEND_NICKNAME"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func DeclineFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.DeclineReqMgo(r.PostFormValue("FRIEND_NICKNAME"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func ConnectFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.SendReqMgo(r.FormValue("FRIEND_NICKNAME"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func DisconnectFriend(w http.ResponseWriter, r *http.Request) {
	nicknameStr := r.FormValue("FRIEND_NICKNAME")
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.RemoveFriendMgo(nicknameStr)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func networkRelation(w http.ResponseWriter, r *http.Request) {
	nicknameStr := pat.Param(r, "nickname")
	ctrl := getController(r)
	defer ctrl.session.Close()
	networkRelation := ctrl.getNetworkRelation(nicknameStr)
	ifErr(json.NewEncoder(w).Encode(networkRelation))
}
