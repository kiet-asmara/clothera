package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "root"
	dbPassword = "bE3eF4EDf5-gacEGEh2E43eB3AEG2daE"
	dbHost     = "viaduct.proxy.rlwy.net"
	dbPort     = "38725"
	dbName     = "railway"
	// local      = "root:@tcp(127.0.0.1:3306)/clothera"
)

func GetDB() (*sql.DB, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/clothera")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
