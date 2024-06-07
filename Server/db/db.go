// db/db.go
package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS usdbrl ( id INTEGER NOT NULL PRIMARY KEY, code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, createDate TEXT)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}
