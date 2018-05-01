package controllers

import (
	"net/http"
)

type Home struct {
	Base
}

func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	h.Render(w, nil, "layout", "home.index")
}

func (h Home) About(w http.ResponseWriter, r *http.Request) {
	h.Render(w, nil, "layout", "home.about")
}
