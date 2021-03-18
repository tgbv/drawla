package lib

import (
	"html/template"
	"path/filepath"
)

/*
*	returns template absolute path
	the implicit path is filepath.Abs("./view/")
*/
func GetTmplPath(n string) string {
	p, err := filepath.Abs("./view/" + n)
	if err != nil {
		panic(err)
	}

	return p
}

/*
*	returns a template.Template instance of the template from path
	the implicit path is filepath.Abs("./view/")
*/
func GetTmpl(n string) *template.Template {
	t, err := template.ParseFiles(GetTmplPath(n))
	if err != nil {
		panic(err)
	}

	return t
}
