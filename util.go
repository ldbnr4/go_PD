package main

import (
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
	for k := range r.PostForm {
		if r.PostForm.Get(k) == "" {
			panic("Bad message")
		}
	}
}
