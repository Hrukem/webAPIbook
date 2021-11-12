package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (env *Store) bookAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	b, err := env.db.GetAll()
	if err != nil {
		log.Println("error get data from database", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	answer, err := json.Marshal(b)
	if err != nil {
		fmt.Println("error Marshal in transport.bookList()", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(answer)
	if err != nil {
		fmt.Println("error Write in transport.bookList()", err)
	}
}
