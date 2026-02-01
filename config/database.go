package config

import (
	"database/sql"
	"log"

	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// test koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// optional: setting pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
