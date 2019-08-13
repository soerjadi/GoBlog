package database

import "database/sql"

type Database interface {
	InitDB() *sql.DB
	InitTestDB() *sql.DB
}

type DBTest interface {
	Clean(tables ...string)
}

type DBTestImpl struct {
	Conn *sql.DB
}
