package transport

import (
	"fmt"
	"golang_ninja/webAPIbook/config"
	"golang_ninja/webAPIbook/pkg/process"
	"golang_ninja/webAPIbook/pkg/storage"
	"log"
	"net/http"
)

type T struct {
	db *storage.DB
	Trnsprt
	process.Proc
}

// Server function start server
func Server() error {
	db, err := storage.NewDb()
	if err != nil {
		log.Println("error create database: ", err)
		return err
	}

	serv := &T{db, Trnsprt{}, process.Proc{}}

	http.HandleFunc("/books", serv.GetAll)
	http.HandleFunc("/book", serv.Post)
	http.HandleFunc("/book/", workDb)

	fmt.Println("start server on port", config.Cfg.Port)
	err = http.ListenAndServe(config.Cfg.Port, nil)
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
