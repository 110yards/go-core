package postgres

import (
	"database/sql"

	"110yards.ca/libs/go/core/logger"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Initialize(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
}

func GetDb() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}
