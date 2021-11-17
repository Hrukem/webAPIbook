package transport

import (
	"golang_ninja/webAPIbook/pkg/process"
	"golang_ninja/webAPIbook/pkg/storage"
	"net/http"
)

type Transport interface {
	Post(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type Trnsprt struct {
	*storage.DB
	Transport
	process.Process
}
