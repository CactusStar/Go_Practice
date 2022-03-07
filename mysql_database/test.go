package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "time"
)


type User struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

type schema struct {
	schemaname   string          `db:"SCHEMA_NAME"`
	DIGEST_TEXT string `db:"DIGEST_TEXT"`
	COUNT_STAR  int            `db:"COUNT_STAR"`
	procsslist string `db:"processlist"`
}

type usercount struct {
	Id int `db:"id"`
	User string `db:"user"`
	Host string `db:"host"`
	db sql.NullString `db:"db"`
	Command string `db:"command"`
	Time string `db:"time"`
	State string `db:"state"`
	Info sql.NullString `db:"Info"`
}

type usercount_count struct {
	count usercount
}
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
	// DB.Exec("USE performance_schema")
	// test, err := DB.Exec("SELECT SCHEMA_NAME,DIGEST_TEXT FROM events_statements_summary_by_digest ORDER BY COUNT_STAR desc LIMIT 1\\G")

	// access most of times table name
	// user := new(schema)
	// row := DB.QueryRow("SELECT SCHEMA_NAME,DIGEST_TEXT, COUNT_STAR FROM events_statements_summary_by_digest ORDER BY COUNT_STAR desc LIMIT 1;")
	// if err := row.Scan(&user.schemaname, &user.DIGEST_TEXT, &user.COUNT_STAR); err != nil {
	// 	fmt.Printf("scan failed, err:%v", err)
	// 	return
	// }
	// usercount2 := new(usercount)
	// row1, err := DB.Query("show processlist;", 1)
	fmt.Println(len(queryMulti(DB)))
	
	// if err := row1.Scan(&usercount2.Id, &usercount2.User, &usercount2.Host, &usercount2.db, &usercount2.Command, &usercount2.Time, &usercount2.State, &usercount2.Info); err != nil {
	// 	fmt.Printf("scan failed1111, err:%v", err)
	// 	return
	// }
	// tablecmd := *&user.DIGEST_TEXT
	// tablename := strings.Split(tablecmd,"`")
	// accesscount := *&user.COUNT_STAR
	// usercount1 := *&usercount2.Id

	// fmt.Println(tablename[1])
	// fmt.Println(accesscount)
	// fmt.Println(usercount1)
	
	// data accessed most of the times

	// // err = DB.Ping()
	// // if err != nil{
	// // 	panic(err.Error())
	// // }
	// DB.SetConnMaxLifetime(100 * time.Second)
	// DB.SetMaxOpenConns(100)
	// DB.SetMaxIdleConns(16)
	// queryOne(DB)
	// queryMulti(DB)
	// insertData(DB)
	// updateData(DB)
	// deleteData(DB)

	// aaa := sql.Drivers()
	// fmt.Println(aaa)
}

// create database
func createDatabase(DB *sql.DB) {
	_, err := DB.Exec("CREATE DATABASE testDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB create successfully")
	}
}

// create table
func createTable() {
	
}

//查询单行
func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}

//查询多行
func queryMulti(DB *sql.DB) []usercount{
	multiprocess := make([]usercount, 0)
	usercount2 := new(usercount)
	rows, err := DB.Query("show processlist;")
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
	}
	for rows.Next() {
		err = rows.Scan(&usercount2.Id, &usercount2.User, &usercount2.Host, &usercount2.db, &usercount2.Command, &usercount2.Time, &usercount2.State, &usercount2.Info)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
		}
		fmt.Print(*usercount2)
		multiprocess = append(multiprocess, *usercount2)
	}
	return multiprocess
}

//插入数据
func insertData(DB *sql.DB){
	result,err := DB.Exec("insert INTO users(name,age) values(?,?)","YDZ",23)
	if err != nil{
		fmt.Printf("Insert failed,err:%v",err)
		return
	}
	lastInsertID,err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v",err)
		return
	}
	fmt.Println("LastInsertID:",lastInsertID)
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}

//更新数据
func updateData(DB *sql.DB){
	result,err := DB.Exec("UPDATE users set age=? where id=?","30",3)
	if err != nil{
		fmt.Printf("Insert failed,err:%v",err)
		return
	}
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}

//删除数据
func deleteData(DB *sql.DB){
	result,err := DB.Exec("delete from users where id=?",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v",err)
		return
	}
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}