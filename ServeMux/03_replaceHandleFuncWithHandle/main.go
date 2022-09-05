package main

import (
"html/template"
"log"
"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "This is index")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Dog Dog Dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Greeting. The name is Kenshi Kuo.")
	if err != nil {
		log.Fatalln(err)
	}
}
