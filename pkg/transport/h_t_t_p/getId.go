package h_t_t_p

import (
	"fmt"
	"log"
	"net/http"
)

// GetID function return object by id
func (t *T) GetID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	answer, err := t.GetObjectFromDb(t.DbPostgres, path, t.LoggingInMongo)
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
