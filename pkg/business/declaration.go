package business

import (
	"golang_ninja/webAPIbook/pkg/storage"
	"net/http"
)

type Business interface {
	InsertObjectInDb(*http.Request, *storage.DB) (map[string]int, error)
	GetAllFromDb(*storage.DB) ([]byte, error)
}

type B struct {
	d1 *storage.DB
	Business
	storage.S
}
