package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type PingerErrorCode int

const (
	connectError PingerErrorCode = iota + 1
	pingError
)

func getDsn() (string, error) {
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		return "", errors.New("missing environment variable: 'DSN' ")
	}
	return dsn, nil
}

func getDbConn() (*sql.DB, error) {
	dsn, err := getDsn()
	if err != nil {
		return nil, err
	}
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func myExit(code PingerErrorCode) {
	fmt.Printf("Aborting application due to exit: %s", code)
	os.Exit(int(code))
}

func main() {
	dbConn, err := getDbConn()
	if  err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		myExit(connectError)
	}
	err = dbConn.Ping()
	if err != nil {
		fmt.Printf("Could not ping database: %s\n", err)
		myExit(pingError)
	}
	fmt.Println("Pinged database successfully")
}


//go:generate stringer -type=PingerErrorCode
