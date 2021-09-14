package database

import (
	"database/sql"
	"fmt"
)

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:salman12@tcp(127.0.0.1:3306)/calculator")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	return db

}
