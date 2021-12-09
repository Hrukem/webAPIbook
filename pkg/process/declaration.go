package process

import (
	"golang_ninja/webAPIbook/pkg/storage"
	"net/http"
)

type Process interface {
	InsertObjectInDb(*http.Request, *storage.DB) (map[string]int, error)
	GetAllFromDb(*storage.DB) ([]byte, error)
}

type Proc struct {
	Process
	storage.S
}
