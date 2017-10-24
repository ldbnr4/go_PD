package main

import (
	"encoding/json"
	"io"
	//"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// AddUserMsg ...
type AddUserMsg struct {
	Username string
	Password string
	Email    string
	Nickname string
}

//AddUserResp ...
type AddUserResp struct {
	Error CreateUserError
	ID    string
}

//DelUserMsg ...
type DelUserMsg struct {
	ID       string
	Password string
}

//LoginMsg ...
type LoginMsg struct {
	Username string
	Password string
}

//UserCreate ...
func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin", "*")

	msg := new(AddUserMsg)
	FillStruct(r, msg)

	insertResp := InsertUser(*msg)
	userDirLocation := PrjDir + insertResp.ID
	ifErr(os.MkdirAll(userDirLocation, os.ModePerm))
	ifErr(copyFileContents(PrjDir+"profile.png", userDirLocation+"/"+insertResp.ID))

	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(insertResp))
}

//UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelUserMsg)
	FillStruct(r, msg)
	RemoveUser(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	msg := new(LoginMsg)
	FillStruct(r, msg)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ifErr(json.NewEncoder(w).Encode(RetrieveUser(*msg)))
}

//ProfPic ...
func ProfPic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UID := vars["UID"]

	f, err := os.Open(PrjDir + UID + "/" + UID)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}
