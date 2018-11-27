package db

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	DB        *sql.DB
	dbRWMutex sync.RWMutex
}

func GetDBConnection(DBName, DNS string) (*sql.DB, error) {

	/*
		user@unix(/path/to/socket)/dbname?charset=utf8
		user:password@tcp(localhost:5555)/dbname?charset=utf8
		user:password@/dbname
		user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
		username:password@tcp(127.0.0.1:3306)/test
	*/

	return sql.Open("mysql", "user:password@/dbname")
}

func (DBObject *Mysql) Close() error {
	err := DBObject.DB.Close()
	return err
}

func (DBObject *Mysql) ExecuteQuery(strQuery string) error {
	DBObject.dbRWMutex.Lock()
	defer DBObject.dbRWMutex.Unlock()

	_, err := DBObject.DB.Exec(strQuery)

	return err
}

func (DBObject *Mysql) SelectQuery(strQuery string) (*sql.Rows, error) {
	DBObject.dbRWMutex.Lock()
	defer DBObject.dbRWMutex.Unlock()

	rows, err := DBObject.DB.Query(strQuery)

	return rows, err
}

func (DBObject *Mysql) UpdateQuery(strQuery, data string) error {
	DBObject.dbRWMutex.Lock()
	defer DBObject.dbRWMutex.Unlock()

	stmt, err1 := DBObject.DB.Prepare(strQuery)

	if err1 != nil {
		return err1
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(data)

	if err2 != nil {
		return err2
	}
	return nil
}
