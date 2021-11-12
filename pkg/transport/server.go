package transport

import (
	"golang_ninja/webAPIbook/pkg/config"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
)

type Store struct {
	db storage.Storage
}

// Server function start server
func Server() {

	db, err := storage.NewDb()
	if err != nil {
		log.Fatal("error create database", err)
	}

	env := &Store{db}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", env.bookAll)
	mux.HandleFunc("/book", env.bookPost)
	mux.HandleFunc("/book/", workDb)

	log.Fatal(http.ListenAndServe(config.Cfg.Port, mux))
}

func workDb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//		old.Get(w, r)
	case "POST":
		//		old.Post(w, r)
	case "PUT":
		//		old.Put(w, r)
	case "DELETE":
		//		old.Delete(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
