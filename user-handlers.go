package main

import (
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	var msg AddUserMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &msg); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//TODO: Call a function that inserts a new User into the mongoDb and returns the Id string
	uid := "<user_id>"
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(uid); err != nil {
		panic(err)
	}

}
