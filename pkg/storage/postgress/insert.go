package postgress

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

func (s *S) Insert(b Book, db *DB) (int, error) {
	id := 0
	str := "INSERT INTO books " +
		"(title, author, publishing, dateinsert) " +
		"values ($1, $2, $3, $4) returning id"
	err1 := db.QueryRow(
		str,
		b.Title,
		b.Author,
		b.Publishing,
		time.Now(),
	).Scan(&id)

	if err1 != nil {
		if err1 == sql.ErrNoRows {
			log.Println("no create rows in storage.PostBook()")
			return 0, err1
		}
	}

	idString := strconv.Itoa(id)

	db.Set(idString, b, 15)

	return id, nil
}
