package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (env *Store) bookPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	var b map[string]string
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println("error decode body in transport.bookCreate()", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := env.db.PostBook(b)
	if err != nil {
		log.Println("error create object in transport.bookCreate()")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error create object in database on server")); err != nil {
			log.Println("error write body answer in transport.bookCreate()", err)
		}
		return
	}

	m := map[string]int{"id": id}
	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in transport.Put()")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error JSON on server")); err != nil {
			log.Println("error write body answer in transport.bookCreate()", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(answer); err != nil {
		log.Println("error write body answer in transport.bookCreate()", err)
	}
}
