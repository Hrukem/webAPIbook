// Package processing implements functions for processing requests
// and working with the database
package processing

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
)

type books struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishing string `json:"publishing"`
}

// GetList function returns a list of all objects in response to the request
func GetList(w http.ResponseWriter, r *http.Request) {
	db, err := storage.ConnectDb()
	if err != nil {
		log.Println("error database connection in processing.Get()", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error database connection")); err != nil {
			log.Println("error write body answer in processing.Get()", err)
		}
		return
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Println("error close db in processing.Get()", err)
		}
	}()

	str := "SELECT id, title, author, publishing FROM books"
	rows, err := db.Query(str)
	if err != nil {
		log.Println("error get data from db in processing.GetList")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error get from database")); err != nil {
			log.Println("error write body answer in processing.GetList()", err)
		}
		return
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Println("error close rows in processing.GetList()")
		}
	}()

	bookList := make([]books, 0)
	for rows.Next() {
		b := books{}
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publishing)
		if err != nil {
			fmt.Println("error Scan rows in processing.GetList()")
			continue
		}
		bookList = append(bookList, b)
	}

	answer, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println("error Marshal in processing.GetList()", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(answer)
	if err != nil {
		fmt.Println("error Write in processing.GetList()", err)
	}
}
