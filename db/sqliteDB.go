package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

func Setup() *SqliteDB {
	db, err := sql.Open("sqlite3", "./decrypted_database.db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &SqliteDB{
		DB: db,
	}
}

func (database *SqliteDB) Close() {
	database.DB.Close()
}
