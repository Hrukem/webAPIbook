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

func GenerationJWT(r *http.Request) ([]byte, int) {
	secret := []byte(config.Cfg.SecretJWT)

	path := r.URL.Path
	pswd := strings.ReplaceAll(path, "/auth/", "")
	if pswd == "" {
		m := map[string]string{"error": "bad request"}
		answer, err := json.Marshal(m)
		if err != nil {
			log.Println("error Marshal in authentication.GenerationJWT()")
		}

		return answer, http.StatusBadRequest
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("error get token in authentication.GenerationJWT() row 46:", err)
		m := map[string]string{"error": "error create token"}
		answer, err := json.Marshal(m)
		if err != nil {
			log.Println("error Marshal in authentication.GenerationJWT()", err)
		}

		return answer, http.StatusInternalServerError
	}

	m := map[string]string{"token": tokenString}
	answer, err := json.Marshal(m)
	if err != nil {
		log.Println("error Marshal in authentication.GenerationJWT()", err)
	}
	return answer, http.StatusCreated
}
