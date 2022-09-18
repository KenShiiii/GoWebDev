package main

import (
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First	string
	Last 	string
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {

}

func bar(w http.ResponseWriter, req *http.Request)  {

}

