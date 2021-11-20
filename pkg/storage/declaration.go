package storage

import (
	"database/sql"
	"github.com/Hrukem/domz2_1"
	_ "github.com/lib/pq"
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/cache"
	"log"
)

type Storage interface {
	SelectAll(*DB) ([]books, error)
	Insert(map[string]string, *DB) (int, error)
}

type DB struct {
	*sql.DB
	*domz2_1.Cache
}
type S struct {
	DB
	Storage
}

type Book struct {
	Title      string
	Author     string
	Publishing string
}

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
