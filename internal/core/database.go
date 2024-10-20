package core

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Connect() {
	DB = sqlx.MustConnect("sqlite3", "./sqlite.db")
}
