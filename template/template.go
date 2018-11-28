package template

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("htmltemplates/*.gohtml"))
}

func Index(w http.ResponseWriter, res *http.Request) {

	gd := struct {
		Title string
	}{Title: "Index Page"}

	tpl.ExecuteTemplate(w, "index.gohtml", gd)
}

func About(w http.ResponseWriter, r *http.Request) {
	gd := struct {
		Title string
	}{Title: "ABOUT"}

	tpl.ExecuteTemplate(w, "about.gohtml", gd)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	gd := struct {
		Title string
	}{Title: "CONTACT"}

	tpl.ExecuteTemplate(w, "contact.gohtml", gd)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	gd := struct {
		Title string
	}{Title: "Signup"}

	tpl.ExecuteTemplate(w, "signup.gohtml", gd)
}

func Process(w http.ResponseWriter, r *http.Request) {

	fn := r.FormValue("first")
	ln := r.FormValue("last")

	type GData struct {
		Title string
	}

	d := struct {
		GData
		First string
		Last  string
	}{
		GData: GData{
			Title: "PROCESS",
		},
		First: fn,
		Last:  ln,
	}

	tpl.ExecuteTemplate(w, "process.gohtml", d)
}
