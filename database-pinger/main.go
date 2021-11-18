package main

// import "alias" "package"
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //
	"os"
)

const (
	connectError = iota + 1
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

func main() {
	dbConn, err := getDbConn()
	if  err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		os.Exit(connectError)
	}
	err = dbConn.Ping()
	if err != nil {
		fmt.Printf("Could not ping database: %s\n", err)
		os.Exit(pingError)
	}
	fmt.Println("Pinged database successfully")
}
