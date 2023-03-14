package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (db *sql.DB) {

	dbDriver := "mysql"

	dbUser := "root"
	dbPass := "root"
	dbName := "crud-example"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}
