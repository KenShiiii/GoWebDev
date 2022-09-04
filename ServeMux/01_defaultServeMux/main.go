package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "This is index")
}

func dog(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "dog dog dog !!")
}

func me(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "The name is Kenshi Kuo")
}
