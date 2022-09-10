package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/Dog.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", "This is from dog")
	if err != nil {
		log.Fatalln(err)
	}
	//io.WriteString(w, `<img src="Dog.jpg" alt="picture not found">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Dog.jpg")
}