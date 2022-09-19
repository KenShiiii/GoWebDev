package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct{
	UserName string
	Password []byte
	First    string
	Last 	 string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req){
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request)  {
	if alreadyLoggedIn(req){
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken?
		if _, ok:= dbUsers[un];ok{
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//	hash the password
		encryptP, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil{
			http.Error(w,"Internal server error", http.StatusInternalServerError)
			return
		}

		//	create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//	store user in dbUsers
		u := user{un,encryptP,f,l}
		dbUsers[un] = u

		u = user{un,encryptP, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}