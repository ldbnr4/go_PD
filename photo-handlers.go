package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"goji.io/pat"
)

func PhotoCreate(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.InsertPhoto(getFile(r), r.PostFormValue("AID"))

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//GetPhoto ...
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	ctrl := getPDUController(r)
	defer ctrl.session.Close()

	w.Header().Set("Cache-Control", "public, max-age=31536000")
	pid := pat.Param(r, "PID")

	f, err := os.Open(PrjDir + ctrl.User.ObjectId.Hex() + "/" + pid)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
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
	ctrl := getPDUController(r)
	defer ctrl.session.Close()
	ctrl.DeletePhoto(r.PostFormValue("PID"))
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
