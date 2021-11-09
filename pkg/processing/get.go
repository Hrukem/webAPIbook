// Package processing implements functions for processing requests
// and working with the database
package processing

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
	"strconv"
)

type book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publishing string `json:"publishing"`
}

// Get function in response to the request,
// it returns a record from the database at the specified id
func Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("error convert string to int in processing.Get()", err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("error id, not integer")); err != nil {
			log.Println("error write body answer in processing.Get()", err)
		}
		return
	}

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

	str := "SELECT title, author, publishing FROM books where id = $1"
	row := db.QueryRow(str, idInt)
	b := book{}
	errSql := row.Scan(
		&b.Title,
		&b.Author,
		&b.Publishing,
	)

	var res interface{}

	if errSql != nil {
		res = make([]string, 0)
	} else {
		res = b
	}

	answer, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error Marshal in processing.Get()", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(answer)
	if err != nil {
		fmt.Println("error Write in processing.Get()", err)
	}
}
