package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func ConnectToDB(connStr string) *sql.DB {
	if db == nil {
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Unable to open connection to db: %w", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Unable to ping db: %w", err)
		}

		db.SetConnMaxLifetime(0)
		db.SetMaxIdleConns(50)
		db.SetMaxOpenConns(50)
	}

	return db
}

func Query(q string, args ...any) (*sql.Rows, error) {
	if db == nil {
		return nil, fmt.Errorf("db is not instantiated")
	}

	return db.Query(q)
}
