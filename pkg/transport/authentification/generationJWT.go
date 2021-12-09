package authentification

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"golang_ninja/webAPIbook/pkg/config"
	"log"
	"net/http"
	"strings"
	"time"
)

//func (t *T) GenerationJWT(w h_t_t_p.ResponseWriter, r *h_t_t_p.Request) {

func GenerationJWT(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pswd := strings.ReplaceAll(path, "/auth/", "")
	if pswd == "" {
		m := map[string]string{"error": "bad request"}
		answer, err := json.Marshal(m)
		if err != nil {
			log.Println("error Marshal in transport.GenerationJWT()")
		}
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(answer)
		if err != nil {
			log.Println("error Write in transport.GenerationJWT()")
		}
		return
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Cfg.SecretJWT))
	if err != nil {
		log.Println("error get token in transport.GenerationJWT() row 46:", err)
		m := map[string]string{"error": "error create token"}
		answer, err := json.Marshal(m)
		if err != nil {
			log.Println("error Marshal in transport.GenerationJWT()", err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(answer)
		if err != nil {
			log.Println("error Write in transport.GenerationJWT()", err)
		}
		return
	}

	m := map[string]string{"token": tokenString}
	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in transport.GenerationJWT()", err)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(answer)
	if err != nil {
		log.Println("error Write in transport.GenerationJWT()", err)
	}
}
