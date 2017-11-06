package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SaveImageFile(file multipart.File, owner, pid string) {
	path := PrjDir + owner + "/" + pid

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	ifErr(err)
	copied, err := io.Copy(f, file)
	ifErr(err)
	fileInfo, err := f.Stat()
	ifErr(err)
	if fileInfo.Size() != copied {
		panic("Did not properly save image!")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Saved file is not on the disk!")
	}
	defer file.Close()
	defer f.Close()
}

func serveFile(profPicPath string, w http.ResponseWriter) {
	if _, err := os.Stat(profPicPath); os.IsNotExist(err) {
		panic("This file does not exist!")
	}
	f, err := os.Open(profPicPath)
	ifErr(err)
	bytesCopied, err := io.Copy(w, f)
	ifErr(err)
	fileInfo, err := f.Stat()
	ifErr(err)
	if fileInfo.Size() != bytesCopied {
		panic("Unable to send back all photo data to the client!")
	}
	defer f.Close()
}

func removePhotoFile(pid, owner string) {
	switch {
	case pid == "":
		panic("RemovePhoto: empty pid")
	case owner == "":
		panic("RemovePhoto: empty owner")
	}
	path := PrjDir + owner + "/" + pid
	ifErr(os.Remove(path))
}

func getFile(r *http.Request) multipart.File {
	// TODO check if file is an image and what type
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	return file
}
