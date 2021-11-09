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

// Post function updates a record in the database by the specified id
func Post(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("error convert string to int in processing.Delete()", err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("error id, not integer")); err != nil {
			log.Println("error write body answer in processing.Delete()", err)
		}
		return
	}
	var b map[string]string
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := storage.ConnectDb()
	if err != nil {
		log.Println("error database connection in processing.Delete()", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error database connection")); err != nil {
			log.Println("error write body answer in processing.Delete()", err)
		}
		return
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Println("error close db in processing.Delete()", err)
		}
	}()

	for k, v := range b {
		_, err = db.Exec(
			"UPDATE books SET "+k+" = $1 WHERE id = $2", v, idInt)
		if err != nil {
			log.Println("error update in processing.Post()", err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successful"))
}
