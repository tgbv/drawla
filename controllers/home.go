package controllers

import (
	"net/http"

	"github.com/tgbv/drawla/lib"
)

type meme struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	//name := pat.Param(r, "name")
	//fmt.Fprintf(w, "Hello! You have reached the homepage")

	t := lib.GetTmpl("home.htm")

	t.Execute(w, &meme{Title: "memez"})
}
