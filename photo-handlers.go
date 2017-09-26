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
	msg := new(AddPhotoMsg)

	r.Header.Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	FillStruct(r, msg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pid := bson.NewObjectId()

	SaveImageFile(pid.Hex(), *msg)
	InsertPhoto(*msg, pid)

	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	ifErr(json.NewEncoder(w).Encode(pid.Hex()))
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "image/jpeg")
	// albumID := r.URL.Query().Get("albumId")
	imageID := r.URL.Query().Get("imageId")
	userID := r.URL.Query().Get("userId")

	f, err := os.Open(PrjDir + userID + "/" + imageID)
	ifErr(err)
	io.Copy(w, f)
}
