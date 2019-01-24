package router

import (
	"database/sql"
)

var (
	dbPool   *sql.DB
)

func Init(db *sql.DB) {
	dbPool = db

}
