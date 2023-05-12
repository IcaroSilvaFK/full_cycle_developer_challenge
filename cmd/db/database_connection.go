package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabaseConnection() *sql.DB {
	// 																:memory:
	conn, err := sql.Open("sqlite3", "dev.db")

	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	initializeTables(conn)

	return conn
}

func initializeTables(db *sql.DB) {
	transaction := `
		CREATE TABLE IF NOT EXISTS quotations (
			id TEXT PRIMARY KEY,
			code TEXT NOT NULL, 
			codein TEXT NOT NULL, 
			name TEXT NOT NULL, 
			high TEXT NOT NULL, 
			low TEXT NOT NULL, 
			varBid TEXT NOT NULL, 
			pctChange TEXT NOT NULL, 
			bid TEXT NOT NULL, 
			ask TEXT NOT NULL, 
			timestamp DATETIME NOT NULL
		);
	`

	stmt, err := db.Prepare(transaction)

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()

	fmt.Println("Execute create table")

}
