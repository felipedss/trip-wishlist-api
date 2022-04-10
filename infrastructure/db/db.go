package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/felipedss/trip-wishlist-api/infrastructure/env"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func StartDB() {

	environment := env.GetEnvConfig().DBConfig

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", environment.User, environment.Password, environment.Host, environment.Database)
	database, err := sql.Open(environment.Driver, connectionString)
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
