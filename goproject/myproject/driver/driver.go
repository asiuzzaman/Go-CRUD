package driver

import (
	"database/sql"
	"fmt"
)

// DB my postgres database
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
	//	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ... take the params from main
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf("host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable", host, port, uname, pass, dbname)
	d, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}
