package storage

import (
	"database/sql"
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/cache"
	"log"
)

func NewDb() (*DB, error) {
	ch := cache.Ch{}
	cch := ch.NewCache()

	str := "host=" + "localhost" + //config.Cfg.Dbhost +
		" port=" + config.Cfg.Dbport +
		" user=" + config.Cfg.Dbuser +
		" password=" + config.Cfg.Dbpass +
		" dbname=" + config.Cfg.Dbname +
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

	return &DB{db, cch}, nil
}
