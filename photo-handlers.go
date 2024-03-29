package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"goji.io/pat"
)

func PhotoCreate(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.InsertPhoto(getFile(r), r.PostFormValue("AID"))

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetPhoto ...
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	uid := pat.Param(r, "UID")
	env := pat.Param(r, "ENV")
	r.Header.Set("UID", uid)
	r.Header.Set("ENV", env)
	ctrl := getController(r)
	defer ctrl.session.Close()

	pidStr := pat.Param(r, "PID")
	pid := convertToObjectID(pidStr)
	photoObj := getPhotoObj(pid, ctrl.photoCol)
	picPath := PrjDir + photoObj.Owner.Hex() + "/" + pidStr
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	serveFile(picPath, w)
}

//DevHero ...
func DevHero(w http.ResponseWriter, r *http.Request) {
	picURL := pat.Param(r, "hero")
	f, err := os.Open(PrjDir + "_heros/" + picURL)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}

func PhotoDelete(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.DeletePhoto(r.PostFormValue("PID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func profPic(w http.ResponseWriter, r *http.Request) {
	nickname := pat.Param(r, "NICKNAME")
	UID := pat.Param(r, "UID")
	ENV := pat.Param(r, "ENV")
	r.Header.Set("UID", UID)
	r.Header.Set("ENV", ENV)

	ctrl := getController(r)
	defer ctrl.session.Close()

	uidStr := ctrl.getUIDFromNickname(nickname).Hex()

	profPicPath := PrjDir + uidStr + "/" + uidStr
	serveFile(profPicPath, w)
}
