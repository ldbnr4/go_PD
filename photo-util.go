package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"
)

func SaveImageFile(pid string, msg AddPhotoMsg) {
	data := strings.TrimPrefix(msg.Data, "data:;base64,")
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	ifErr(err)

	path := PrjDir + msg.Owner + "/" + pid
	ifErr(ioutil.WriteFile(path, dataBytes, 0644))
}

func RemovePhoto(pid string, owner string) {
	//TODO: implement this
	switch {
	case pid == "":
		panic("RemovePhoto: empty pid")
	case owner == "":
		panic("RemovePhoto: empty owner")
	}
	path := PrjDir + owner + "/" + pid
	ifErr(os.Remove(path))
}
