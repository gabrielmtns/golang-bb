package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	service "go-bb-test/src/service/usecases"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware (h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == ""  {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var authorizationHeaderParser = strings.Split(authorizationHeader, " ")

		if authorizationHeaderParser[0] != "Bearer" || authorizationHeaderParser[1] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}


		fmt.Println(authorizationHeaderParser[1])
		token, err := jwt.ParseWithClaims(authorizationHeaderParser[1], &service.TokenStandard{}, func(token *jwt.Token)(interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		validToken, ok := token.Claims.(*service.TokenStandard)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := validToken.Valid(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}