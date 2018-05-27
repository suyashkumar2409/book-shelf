package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseHandler struct{
	db * sql.DB
}

const(
	dbName = "db/dev.db"
)

func NewDatabaseHandler() *DatabaseHandler {
	return &DatabaseHandler{}
}

func (db *DatabaseHandler)LoadDB() error{
	var err error
	db.db, err = sql.Open("sqlite3", dbName)
	return err
}

func (db *DatabaseHandler)IsAlive() bool{
	return db.db.Ping() == nil
}