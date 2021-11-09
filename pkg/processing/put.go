// Package processing implements functions for processing requests
// and working with the database
package processing

import (
	"encoding/json"
	"fmt"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
	"time"
)

// Put function creates a new entry in the database
// returns the id of the created object
func Put(w http.ResponseWriter, r *http.Request) {
	var b map[string]string
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in processing.Put()", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := storage.ConnectDb()
	if err != nil {
		log.Println("error database connection in processing.Put()", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error database connection")); err != nil {
			log.Println("error write body answer in processing.Put()", err)
		}
		return
	}

	id := 0
	str := "INSERT INTO books " +
		"(author, title, publishing, dateinsert) " +
		"values ($1, $2, $3, $4) returning id"
	db.QueryRow(
		str,
		b["author"],
		b["title"],
		b["publishing"],
		time.Now(),
	).Scan(&id)
	defer func() {
		if err = db.Close(); err != nil {
			log.Println("error close db in processing.Put()")
		}
	}()
	if err != nil {
		log.Println("error database insert PUT", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error database insert")); err != nil {
			log.Println("error write body answer in processing.Put()", err)
		}
		return
	}

	m := map[string]int{"id": id}
	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in processing.Put()")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error JSON on server")); err != nil {
			log.Println("error write body answer in processing.Put()", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(answer); err != nil {
		log.Println("error write body answer in processing.Put()", err)
	}
}
