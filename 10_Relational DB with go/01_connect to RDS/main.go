package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "admin:password@tcp(database-2.cfgevppjdhs4.ap-northeast-1.rds.amazonaws.com:3306)/test_schema?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "successful connect to RDS")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
