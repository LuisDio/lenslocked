package controllers

import (
	"net/http"

	"github.com/LuisDio/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// We need a virtual render
	u.Templates.New.Execute(w, nil)
}
