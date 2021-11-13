package transport

import (
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/business"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
)

type T struct {
	db *storage.DB
	Trnsprt
	business.B
}

// Server function start server
func Server() error {
	db, err := storage.NewDb()
	if err != nil {
		log.Println("error create database: ", err)
		return err
	}

	serv := &T{db, Trnsprt{}, business.B{}}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", serv.GetAll)
	mux.HandleFunc("/book", serv.Post)
	mux.HandleFunc("/book/", workDb)

	err = http.ListenAndServe(config.Cfg.Port, mux)
	if err != nil {
		log.Println("error server", err)
	}
	return nil
}

func workDb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//		old.Get(w, r)
	case "PUT":
		//		old.Put(w, r)
	case "DELETE":
		//		old.Delete(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
