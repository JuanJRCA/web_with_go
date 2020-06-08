package controllers

import (
	"github.com/web_go/views"
)

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {

	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}
