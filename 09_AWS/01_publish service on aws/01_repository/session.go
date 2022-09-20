package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	//	refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//	refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	s := dbSessions[c.Value]
	s.lastActivity = time.Now()
	_, ok := dbUsers[s.un]

	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") //	for demonstration purposes
	showSessions()              //	for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") //	for demonstration purposes
	showSessions()             //	for demonstration purposes
}

//	for demonstration purposes
func showSessions() {
	fmt.Println("**************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
