package main

import (
	"html/template"
	"net/http"
)
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w, nil)
}