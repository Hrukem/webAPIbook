package storage

import (
	"database/sql"
	"golang_ninja/webAPIbook/pkg/process"
	"log"
	"time"
)

func (s S) Insert(book process.Book, db *DB) (int, error) {
	id := 0
	str := "INSERT INTO books " +
		"(author, title, publishing, dateinsert) " +
		"values ($1, $2, $3, $4) returning id"
	err1 := db.QueryRow(
		str,
		book.Title,
		book.Author,
		book.Publishing,
		time.Now(),
	).Scan(&id)

	if err1 != nil {
		if err1 == sql.ErrNoRows {
			log.Println("no create rows in storage.PostBook()")
			return 0, err1
		}
	}

	return id, nil
}
