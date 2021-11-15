package transport

import (
	"log"
	"net/http"
)

func (t *T) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	answer, err := t.InsertObjectInDb(r, t.db)
	if err != nil {
		log.Println("error insert data in database", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error insert object in database on server")); err != nil {
			log.Println("error write body answer in transport.bookCreate()", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(answer); err != nil {
		log.Println("error write body answer in transport.bookCreate()", err)
	}
}
