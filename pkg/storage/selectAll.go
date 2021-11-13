package storage

import "log"

type books struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishing string `json:"publishing"`
}

func (s *S) SelectAll(db *DB) ([]books, error) {
	str := "SELECT id, title, author," + " publishing FROM books"
	rows, err := db.Query(str)
	if err != nil {
		log.Println("error get data from db in interfaces.selectAll()")
		return nil, err
	}

	bookList := make([]books, 0)
	for rows.Next() {
		b := books{}
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publishing)
		if err != nil {
			log.Println("error Scan rows in interfaces.selectAll()")
			return nil, err
		}
		bookList = append(bookList, b)
	}
	return bookList, nil
}
