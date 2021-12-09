package process

import (
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"net/http"
)

type Process interface {
	InsertObjectInDb(*http.Request, *postgress.DB) (map[string]int, error)
	GetAllFromDb(*postgress.DB) ([]byte, error)
	GetObjectFromDb(*postgress.DB, string) ([]byte, error)
}

type Proc struct {
	Process
	postgress.S
}
