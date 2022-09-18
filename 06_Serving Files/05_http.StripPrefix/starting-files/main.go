package main

import (
	"html/template"
	"net/http"
)

func main()  {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", index)
	http.Handle("/public/",http.StripPrefix("/public", fs))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl.Execute(w, nil)
}