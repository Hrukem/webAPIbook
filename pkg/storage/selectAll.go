package storage

import "log"

func (s *S) SelectAll(db *DB) ([]books, error) {
	str := "SELECT id, title, author," + " publishing FROM books"
	rows, err := db.Query(str)
	if err != nil {
		log.Println("error get data from db in storage.SelectAll()")
		return nil, err
	}

	objectList := make([]books, 0)
	for rows.Next() {
		b := books{}
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publishing)
		if err != nil {
			log.Println("error Scan rows in storage.SelectAll()")
			return nil, err
		}
		objectList = append(objectList, b)
	}
	return objectList, nil
}
