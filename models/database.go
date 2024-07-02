package models

import (
    "database/sql"
     _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// SetDB mengatur koneksi database global
func SetDB(database *sql.DB) {
    db = database
}

// GetDB mengembalikan koneksi database global
func GetDB() *sql.DB {
    return db
}
