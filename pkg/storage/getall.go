package storage

import (
	"log"
)

func (db *DB) GetAll() ([]books, error) {
	str := "SELECT id, title, author," + " publishing FROM books"
	rows, err := db.Query(str)
	if err != nil {
		log.Println("error get data from db in transport.GetAll()")
		return nil, err
	}

	bookList := make([]books, 0)
	for rows.Next() {
		b := books{}
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publishing)
		if err != nil {
			log.Println("error Scan rows in transport.GetAll()")
			return nil, err
		}
		bookList = append(bookList, b)
	}
	return bookList, nil
}
