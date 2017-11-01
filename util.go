package main

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/schema"
)

func ifErr(operation error) {
	if operation != nil {
		panic(operation)
	}
}

func FillStruct(r *http.Request, strct interface{}) {
	var decoder = schema.NewDecoder()
	ifErr(r.ParseForm())
	ifErr(decoder.Decode(strct, r.PostForm))
}

func getFile(r *http.Request) multipart.File {
	// TODO check if file is an image and what type
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	return file
}
