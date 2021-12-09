package storage

import "log"

func (s *S) SelectObject(db *DB, id int) (Book, error) {
	str := "SELECT title, author," + " publishing FROM books" + " where id = $1"
	row := db.QueryRow(str, id)

	object := Book{}
	err := row.Scan(&object.Title, &object.Author, &object.Publishing)
	if err != nil {
		log.Println("error Scan row in storage.SelectObject()")
		return Book{}, err
	}

	return object, nil
}
