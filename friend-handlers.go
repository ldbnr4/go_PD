package main

import (
	"encoding/json"
	"net/http"
)

func AcceptFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.AcceptReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func DeclineFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.DeclineReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func ConnectFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.SendReqMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func DisconnectFriend(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.RemoveFriendMgo(r.PostFormValue("FUID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
