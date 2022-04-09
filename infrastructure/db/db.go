package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func StartDB() {
	//database, err := sql.Open("mysql", "mysql:mysql@tcp(content_db:3306)/trip_wishlist_db")
	database, err := sql.Open("mysql", "mysql:mysql@tcp(localhost:3306)/trip_wishlist_db")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	db = database
}

func GetDB() *sql.DB {
	return db
}
