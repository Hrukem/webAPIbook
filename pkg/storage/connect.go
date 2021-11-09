// Package storage implements the database connection function
package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // for use behavior
	"golang_ninja/webAPIbook/pkg/config"
)

// ConnectDb function creates a database connection
func ConnectDb() (*sql.DB, error) {
	str := "host=" + config.Cfg.DbHost +
		" port=" + config.Cfg.DbPort +
		" user=" + config.Cfg.DbUser +
		" password=" + config.Cfg.DbPass +
		" dbname=" + config.Cfg.DbName +
		" sslmode=disable"

	db, err := sql.Open("postgres", str)
	if err != nil {
		fmt.Println("error connect to database")
		return nil, err
	}

	return db, nil
}
