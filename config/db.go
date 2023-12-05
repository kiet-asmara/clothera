package config

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() {
	const (
		dbUser     = "root"
		dbPassword = "a341CgeH16HbaGAHBdBfB3BdHagGggEg"
		dbHost     = "roundhouse.proxy.rlwy.net"
		dbPort     = "38992"
		dbName     = "railway"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
}
