package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	USERNAME = "root"
	PASSWORD = "StayReal1988@"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "testgo"
)

type schema struct {
	schemaname   string          `db:"SCHEMA_NAME"`
	DIGEST_TEXT string `db:"DIGEST_TEXT"`
	COUNT_STAR  int            `db:"COUNT_STAR"`
	procsslist string `db:"processlist"`
}

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	monitorSentence(DB)
}

// max table
func monitorSentence(DB *sql.DB) {
	result := new(schema)
	row := DB.QueryRow("select SCHEMA_NAME,digest_text,COUNT_STAR from performance_schema.events_statements_summary_by_digest ORDER BY COUNT_STAR desc limit 1;")
	if err := row.Scan(&result.schemaname,&result.DIGEST_TEXT, &result.COUNT_STAR); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	fmt.Println(*&result.DIGEST_TEXT)
	fmt.Println(*&result.schemaname)
	tablecmd := *&result.DIGEST_TEXT
	tablename := strings.Split(tablecmd," `")
	fmt.Println(tablename[0])
	fmt.Println(tablename[1])
	accesscount := *&result.COUNT_STAR
	fmt.Println(accesscount)
}
// connected accounts

// generate chart

// web application display