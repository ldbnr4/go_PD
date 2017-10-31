package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveImageFile(pid, owner string, r *http.Request) *AddPhotoMsg {
	r.ParseMultipartForm(32 << 20)
	msg := new(AddPhotoMsg)
	FillStruct(r, msg)

	path := PrjDir + owner + "/" + pid
	// panic(r.PostForm)
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	io.Copy(f, file)
	return msg
}

func FSRemovePhoto(pid, owner string) {
	switch {
	case pid == "":
		panic("RemovePhoto: empty pid")
	case owner == "":
		panic("RemovePhoto: empty owner")
	}
	path := PrjDir + owner + "/" + pid
	ifErr(os.Remove(path))
}
