package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establish connection to db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@/yalla_go")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// Close closed connection of db
func Close(db *sql.DB) error {
	return db.Close()
}
