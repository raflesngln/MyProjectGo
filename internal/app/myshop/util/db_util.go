package util

import (
	"database/sql"
	"os"
)

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("pos"))
	PanicError(err)
	return db
}