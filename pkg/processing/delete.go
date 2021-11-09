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

// Delete function deletes an entry by the specified id from the database
func Delete(w http.ResponseWriter, r *http.Request) {
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

	var str = "DELETE FROM books where id = $1"
	_, err = db.Exec(str, idInt)
	if err != nil {
		fmt.Println("error delete from db in processing.Delete()", err)
	}

	s := "delete successful"
	answer, err := json.Marshal(s)
	if err != nil {
		log.Println("error Marshal in processing.Delete()")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error JSON on server")); err != nil {
			log.Println("error write body answer in processing.Delete()", err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(answer); err != nil {
		log.Println("error write body answer in processing.Put()", err)
	}
}
