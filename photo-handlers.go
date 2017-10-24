package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

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

//DevPhoto ...
func DevPhoto(w http.ResponseWriter, r *http.Request) {
	pics := []string{"ironMan.png", "thor.png", "hulk.png", "spiderMan.png"}
	s := rand.NewSource(time.Now().Unix())
	random := rand.New(s) // initialize local pseudorandom generator
	picURL := pics[random.Intn(len(pics))]
	f, err := os.Open(PrjDir + picURL)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}

func PhotoDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelPhotoMsg)
	FillStruct(r, msg)
	DeletePhoto(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}
