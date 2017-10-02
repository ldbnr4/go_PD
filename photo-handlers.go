package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"
)

/***********
 * STRUCTS *
 ***********/
type AddPhotoMsg struct {
	Album string
	Owner string
	Data  string
}

/************
 * HANDLERS *
 ************/
func PhotoCreate(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pid := bson.NewObjectId()

	msg := SaveImageFile(pid.Hex(), r)

	InsertPhoto(*msg, pid)

	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(pid.Hex()))
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "image/jpeg")
	// albumID := r.URL.Query().Get("albumId")
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	imageID := r.URL.Query().Get("imageId")
	userID := r.URL.Query().Get("userId")

	f, err := os.Open(PrjDir + userID + "/" + imageID)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}
