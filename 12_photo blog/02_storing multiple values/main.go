package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"strings"
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
	cookie := getCookie(w, r)

	cookie = checkValue(w, r)

	xs := strings.Split(cookie.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()

		cookie = &http.Cookie{Name: "session", Value: sid.String()}
		http.SetCookie(w, cookie)
	}

	return cookie
}

func checkValue(w http.ResponseWriter, r *http.Request) *http.Cookie {
	p1 := "dog.jpeg"
	p2 := "cat.jpeg"
	p3 := "beach.jpeg"
	c := getCookie(w, r)

	if !strings.Contains(c.Value, p1) {
		c.Value += "|" + p1
		fmt.Println("p1 added")
	}
	if !strings.Contains(c.Value, p2) {
		c.Value += "|" + p2
		fmt.Println("p2 added")
	}
	if !strings.Contains(c.Value, p3) {
		c.Value += "|" + p3
		fmt.Println("p3 added")
	}

	http.SetCookie(w,c)

	return c
}
