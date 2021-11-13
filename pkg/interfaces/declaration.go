package interfaces

import (
	"database/sql"
	"golang_ninja/webAPIbook/pkg/config"
	"log"
	"net/http"
)

type Transport interface {
	Post(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type Business interface {
	newObjectInDb(*http.Request) (map[string]int, error)
	getAllFromDb()
}

type Storage interface {
	insert()
	selectAll()
}

type DB struct {
	*sql.DB
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