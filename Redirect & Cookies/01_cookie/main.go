package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var n int

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	visitTime(w, req)

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	visitTime(w, req)

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

// Using cookies, track how many times a user has been to your website domain.
func visitTime(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)
}
