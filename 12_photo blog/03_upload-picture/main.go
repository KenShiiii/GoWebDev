package main

import (
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	//	process form submission
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		//	create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		//	create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		//	copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		//	add filename to this user's cookie

		cookie = appendValue(w, cookie, fname)

	}

	xs := strings.Split(cookie.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()

		cookie = &http.Cookie{Name: "session", Value: sID.String()}
		http.SetCookie(w, cookie)
	}

	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	if !strings.Contains(c.Value, fname) {
		c.Value += "|" + fname
	}
	http.SetCookie(w, c)

	return c
}
