package process

import (
	"golang_ninja/webAPIbook/pkg/storage"
	"net/http"
)

type Book struct {
	Title      string
	Author     string
	Publishing string
}

type Business interface {
	InsertObjectInDb(*http.Request, *storage.DB) (map[string]int, error)
	GetAllFromDb(*storage.DB) ([]byte, error)
}

type B struct {
	Business
	storage.S
}
