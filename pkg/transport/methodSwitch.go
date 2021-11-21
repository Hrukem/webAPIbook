package transport

import "net/http"

func (t *T) MethodSwitch(w http.ResponseWriter, r *http.Request) {
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
