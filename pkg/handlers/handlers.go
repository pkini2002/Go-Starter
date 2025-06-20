package handlers

import (
	"net/http"
	"github.com/pkini2002/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.gohtml")
}
