package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // here
	"os"
)


type DbHandler struct {
	Connection *sql.DB
}


func Handler()(*sql.DB, error){
	databaseConfig := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),

	)
	connection, err := sql.Open(os.Getenv("DB_DRIVER"), databaseConfig)
	if err != nil {
		return nil, err
	}

	err = connection.Ping()

	if err != nil {
		return nil, err
	}

	return connection, nil

}

func (db *DbHandler) Query(query string, queryValues ...interface{})(*sql.Rows, error) {
	queryResult, err := db.Connection.Query(query, queryValues...)

	if err != nil {
		return nil, err
	}


	return queryResult, nil
}

func (db *DbHandler) Exec(query string, queryValues ...interface{})(sql.Result, error){
	result, err := db.Connection.Exec(query, queryValues...)

	if err != nil {
		return nil, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func close(conn *sql.DB) error {
	return conn.Close()
}