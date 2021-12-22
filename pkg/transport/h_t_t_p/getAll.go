package h_t_t_p

import (
	"fmt"
	"log"
	"net/http"
)

// @Summary GetAll
// @Description return list books
// @Accept json
// @Produce json
// @Success 200 {json}
// @Failure 400, 404 {json}
// @Router /books [get]

func (t *Tr) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	answer, err := t.GetAllFromDb(t.DbPostgres, t.LoggingInMongo)
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
