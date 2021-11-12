package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang_ninja/webAPIbook/pkg/config"
	"log"
)

type DB struct {
	*sql.DB
}

type Storage interface {
	GetAll() ([]books, error)
	//	getId(id int) (book, error)
	PostBook(m map[string]string) (int, error)
	//	put(id int) error
	//	delete(id int) error
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
