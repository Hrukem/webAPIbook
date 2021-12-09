package h_t_t_p

import (
	"golang_ninja/webAPIbook/pkg/process"
	"golang_ninja/webAPIbook/pkg/storage/postgress"
	"net/http"
)

type Transport interface {
	Post(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetID(w http.ResponseWriter, r *http.Request)
	MethodSwitch(w http.ResponseWriter, r *http.Request)
	//	GenerationJWT(w h_t_t_p.ResponseWriter, r *h_t_t_p.Request)
}

type Trnsprt struct {
	*postgress.DB
	Transport
	process.Process
}
