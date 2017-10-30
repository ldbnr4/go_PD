package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"goji.io/pat"
	"gopkg.in/mgo.v2/bson"
)

/***********
 * STRUCTS *
 ***********/
type AddPhotoMsg struct {
	Album string
	Owner string
}

type DelPhotoMsg struct {
	UID string
	PID string
}

/************
 * HANDLERS *
 ************/
func PhotoCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pid := bson.NewObjectId()

	msg := SaveImageFile(pid.Hex(), r)

	ctrl := getController()
	defer ctrl.session.Close()
	ctrl.InsertPhoto(*msg, pid)

	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(pid.Hex()))
}

//GetPhoto ...
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	imageID := r.URL.Query().Get("imageId")
	userID := r.URL.Query().Get("userId")

	f, err := os.Open(PrjDir + userID + "/" + imageID)
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
	msg := new(DelPhotoMsg)
	FillStruct(r, msg)
	ctrl := getController()
	defer ctrl.session.Close()
	ctrl.DeletePhoto(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
