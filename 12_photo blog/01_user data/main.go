package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()

		cookie = &http.Cookie{Name: "session", Value: sid.String()}
		http.SetCookie(w, cookie)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
