package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
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
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":80", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "at INDEX")
	check(err)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	// query
	rows, err := db.Query("SELECT aName FROM amigos;")
	check(err)
	defer rows.Close()

	//	data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	//	query through each row
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	instance(w,r)
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("CREATE TABLE customer (name VARCHAR(20));")
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	instance(w,r)
	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	instance(w,r)
	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
	instance(w,r)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name ="James";`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	instance(w,r)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	instance(w,r)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

	instance(w,r)
}

func ping(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil{
		check(err)
		return
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	io.WriteString(w, string(bs)+"\n\n\n")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
