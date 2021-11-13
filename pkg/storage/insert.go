package storage

import (
	"database/sql"
	"log"
	"time"
)

func (s S) Insert(m map[string]string, db *DB) (int, error) {
	id := 0
	str := "INSERT INTO books " +
		"(author, title, publishing, dateinsert) " +
		"values ($1, $2, $3, $4) returning id"
	err1 := db.QueryRow(
		str,
		m["author"],
		m["title"],
		m["publishing"],
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
