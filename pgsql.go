// pgsql.go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER   = "doseok"
	DB_PASSWD = "7795004"
	//DB_HOST   = "homeny.madang.com"
	DB_HOST = "172.16.0.235"
	DB_NAME = "pypos"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

var db *sql.DB

func InitDB() {
	var err error
	connstr := fmt.Sprintf("postgres://%s:%s@%s/%s", DB_USER, DB_PASSWD, DB_HOST, DB_NAME)
	//fmt.Println(connstr)
	db, err = sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
}

func CloseDB() {
	db.Close()

}
