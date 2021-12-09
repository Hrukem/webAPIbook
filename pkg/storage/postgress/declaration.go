package postgress

import (
	"database/sql"
	"github.com/Hrukem/domz2_1"
	_ "github.com/lib/pq"
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

type books struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishing string `json:"publishing"`
}

type Book struct {
	Title      string
	Author     string
	Publishing string
}
