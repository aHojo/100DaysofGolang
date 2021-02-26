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
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/gosql?charset=utf8")
	check(err)
	defer db.Close()

	if err = db.Ping(); err != nil {
		check(err)
	}

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "I made a connection to the db")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
