package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func SaveImageFile(file multipart.File, owner, pid string) {
	path := PrjDir + owner + "/" + pid

	defer file.Close()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	io.Copy(f, file)
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
