package h_t_t_p

import (
	"net/http"
)

func (t *Tr) MethodSwitch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t.GetID(w, r)
	case "PUT":
		//		old.Put(w, r)
	case "DELETE":
		//		old.Delete(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
