package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const BatchSize = 10_000

type DB struct {
	DB *sql.DB
}

func NewDB(dbEnable bool, user, password, host string) (*DB, error) {
	if !dbEnable {
		return nil, nil
	}

	dbSrc := fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host)
	dbConnect, err := sql.Open("mysql", dbSrc)
	if err != nil {
		return nil, err
	}

	return &DB{DB: dbConnect}, nil
}

func (d *DB) Query(queryString string) (*sql.Rows, error) {
	res, err := d.DB.Query(queryString)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *DB) Execute(queryString string, data []interface{}) (sql.Result, error) {
	res, err := d.DB.Exec(queryString, data...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
