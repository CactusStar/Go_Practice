package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "strconv"
)
const (
	USERNAME = "root"
	PASSWORD = "StayReal1988@"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "performance_schema"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	insertquery := "INSERT INTO testgo.test3 (name,age) VALUES"
	hugeaction(5000, DB, insertquery)
}

// huge insert
func hugeaction(count int, DB *sql.DB, query string) {
	for i := 0; i < count; i ++ {
		// tempquery := query + "(" + strconv.Itoa(i) + "," + "'hahaname'" + "," + "'hahaage'" + ")"
		// fmt.Println(tempquery)
		_,err := DB.Exec(query + "(" + "'hahaname'" + "," + "'hahaage'" + ")")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// multiple usesr/thread insert
func multipleaction(count int, DB *sql.DB, query string) {
	for i := 0; i < count; i ++ {
		fmt.Println(i)
		go DB.Exec(query + "(" + "'zzname'" + "," + "'zzage'" + ")")
	}
}
