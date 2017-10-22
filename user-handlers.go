package main

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"os"
)

/***********
 * STRUCTS *
 ***********/
type AddUserMsg struct {
	Username string
	Password string
	Email    string
	Nickname string
}

type AddUserResp struct {
	Id string
}

type DelUserMsg struct {
	Id       string
	Password string
}

type LoginMsg struct {
	Username string
	Password string
}

/************
 * HANDLERS *
 ************/
func UserCreate(w http.ResponseWriter, r *http.Request) {

	msg := new(AddUserMsg)

	FillStruct(r, msg)

	uid := InsertUser(*msg).Hex()
	ifErr(os.MkdirAll(PrjDir+uid, os.ModePerm))

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(AddUserResp{uid}))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelUserMsg)
	FillStruct(r, msg)
	RemoveUser(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	msg := new(LoginMsg)
	FillStruct(r, msg)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ifErr(json.NewEncoder(w).Encode(RetrieveUser(*msg)))
}
