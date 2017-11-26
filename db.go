package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://rubak:rubak@localhost/rubak_test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
func main() {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM rubak1")
	checkErr(err)
	fmt.Println("uid | username | email ")
	for rows.Next() {
		var uid int
		var username string
		var email string
		err = rows.Scan(&uid, &username, &email)
		checkErr(err)
		fmt.Println( uid, username, email)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
