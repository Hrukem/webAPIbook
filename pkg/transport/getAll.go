package transport

import (
	"fmt"
	"log"
	"net/http"
)

func (t *T) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	answer, err := t.GetAllFromDb(t.db)
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
