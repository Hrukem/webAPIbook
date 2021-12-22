package h_t_t_p

import (
	"log"
	"net/http"
)

func (t *Tr) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	answer, err := t.InsertObjectInDb(r, t.DbPostgres, t.LoggingInMongo)
	if err != nil {
		log.Println("error insert data in database", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("error insert object in database on grpcServer")); err != nil {
			log.Println("error write body answer in transportTest.bookCreate()", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(answer); err != nil {
		log.Println("error write body answer in transportTest.bookCreate()", err)
	}
}
