package authentification

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang_ninja/webAPIbook/pkg/config"
	"log"
	"net/http"
	"strings"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("not authenticated"))
			if err != nil {
				fmt.Println("error Write in grpcServer", err)
			}
			return
		}
		// Аутентификация прошла успешно, направляем запрос следующему обработчику
		// next.ServeHTTP(w, r)
		next(w, r)
	}
}

func isAuthenticated(r *http.Request) bool {
	a := r.Header.Get("Authorization")
	sliceStrings := strings.Split(a, " ")

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(sliceStrings[1], claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.SecretJWT), nil
	})

	if err != nil {
		log.Println("error ParseWithClaims: ", err)
		return false
	}
	if !token.Valid {
		log.Println("token not valid: ", token)
		return false
	}
	return true
}
