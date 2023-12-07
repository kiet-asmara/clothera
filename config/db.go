package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "root"
	dbPassword = "a341CgeH16HbaGAHBdBfB3BdHagGggEg"
	dbHost     = "roundhouse.proxy.rlwy.net"
	dbPort     = "38992"
	dbName     = "railway"
	// local      = "root:@tcp(127.0.0.1:3306)/clothera"
)

func GetDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
