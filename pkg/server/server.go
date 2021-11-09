package server

import (
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/processing"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/book", workDb)
	mux.HandleFunc("/books", processing.GetList)

	log.Fatal(http.ListenAndServe(config.Cfg.Port, mux))
}

func workDb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		processing.Get(w, r)
	case "POST":
		processing.Post(w, r)
	case "PUT":
		processing.Put(w, r)
	case "DELETE":
		processing.Delete(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
