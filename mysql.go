package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select id,login_time from user where id=1")
	checkErr(err)
	for rows.Next() {
		var id int
		var lastLoginTime string
		err = rows.Scan(&id, &lastLoginTime)
		checkErr(err)
		DefaultTimeLoc := time.Local
		loginTime, err := time.ParseInLocation("2006-01-02 15:04:05", lastLoginTime, DefaultTimeLoc)
		checkErr(err)
		fmt.Println(loginTime)
	}
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
