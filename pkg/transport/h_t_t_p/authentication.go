package h_t_t_p

import (
	"golang_ninja/webAPIbook/pkg/process/authentification"
	"log"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request) {
	answer, status := authentification.GenerationJWT(r)

	w.WriteHeader(status)
	_, err := w.Write(answer)
	if err != nil {
		log.Println("error Write in transportTest.GenerationJWT()")
	}
}
