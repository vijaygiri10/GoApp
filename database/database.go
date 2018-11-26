package db

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/mutecomm/go-sqlcipher"
)

type Sqlitedb struct {
	DB        *sql.DB
	dbRWMutex sync.RWMutex
}

func GetDBConnection(DBName, Schema string) (*Sqlitedb, error) {

	var Error error
	dbConn.DB, Error = sql.Open("mysql", DBName)

	if Error != nil {
		return nil, Error
	}
	_, err := dbConn.DB.Exec(Schema)

	if err != nil {
		err_close := dbConn.Close()
		if err_close != nil {
			return dbConn, errors.New(err.Error() + err_close.Error())
		} else {
			return dbConn, err
		}
	}

	return dbConn, nil
}

func (DBObject *Sqlitedb) Close() error {
	err := DBObject.DB.Close()
	return err
}

func (DBObject *Sqlitedb) ExecuteQuery(strQuery string) error {
	DBObject.dbRWMutex.Lock()
	defer DBObject.dbRWMutex.Unlock()

	_, err := DBObject.DB.Exec(strQuery)

	return err
}

func (DBObject *Sqlitedb) SelectQuery(strQuery string) (*sql.Rows, error) {
	DBObject.dbRWMutex.Lock()
	defer DBObject.dbRWMutex.Unlock()

	rows, err := DBObject.DB.Query(strQuery)

	return rows, err
}

func (DBObject *Sqlitedb) UpdateQuery(strQuery, data string) error {
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
