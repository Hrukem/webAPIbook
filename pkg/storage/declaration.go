package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang_ninja/webAPIbook/config"
	"log"
)

type Storage interface {
	SelectAll(*DB) ([]books, error)
	Insert(map[string]string, *DB) (int, error)
}

type DB struct {
	*sql.DB
}

type S struct {
	DB
	Storage
}

func NewDb() (*DB, error) {
	str := "host=" + config.Cfg.DbHost +
		" port=" + config.Cfg.DbPort +
		" user=" + config.Cfg.DbUser +
		" password=" + config.Cfg.DbPass +
		" dbname=" + config.Cfg.DbName +
		" sslmode=disable"

	db, err := sql.Open("postgres", str)
	if err != nil {
		log.Println("error connect to database")
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("error Ping to database")
		return nil, err
	}

	return &DB{db}, nil
}
