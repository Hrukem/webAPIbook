package transport

import (
	"fmt"
	"log"
	"net/http"
)

func (t *T) GetID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	answer, err := t.GetObjectFromDb(t.DbServer, path)
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
