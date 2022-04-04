package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Pool *sql.DB

func init() {
	log.Println("initializing mysql")
	cfg := mysql.Config{
		User:                 "testservice",
		Passwd:               "Service321",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	var err error
	Pool, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Pool.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

}
