package main

import (
	"net/http"

	"github.com/gorilla/schema"
)

func FillStruct(r *http.Request, strct interface{}) {
	var decoder = schema.NewDecoder()
	ifErr(r.ParseForm())
	ifErr(decoder.Decode(strct, r.PostForm))
}
