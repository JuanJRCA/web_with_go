package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func Parse(r *http.Request, data interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	dec := schema.NewDecoder()
	err = dec.Decode(data, r.PostForm)

	if err != nil {
		return err
	}

	return nil

}
