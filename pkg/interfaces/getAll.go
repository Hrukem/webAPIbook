package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

func (db *interfaces.DB) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	answer, err := db.getAllFromDb()
	if err != nil {
		log.Println("error get data from database", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(answer)
	if err != nil {
		fmt.Println("error Write in interfaces.GetAll()", err)
	}
}