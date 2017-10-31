package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"goji.io/pat"
	"gopkg.in/mgo.v2/bson"
)

func PhotoCreate(w http.ResponseWriter, r *http.Request) {
	ctrl := getController(r)
	defer ctrl.session.Close()

	pid := bson.NewObjectId()

	msg := SaveImageFile(pid.Hex(), ctrl.UID.Hex(), r)

	ctrl.InsertPhoto(*msg, pid)

	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	ctrl := getController(r)
	defer ctrl.session.Close()
	ctrl.DeletePhoto(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
