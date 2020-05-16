package middleware

import (
	"net/http"

	"github.com/jgromero7/gos-twitter/database"
	jwtService "github.com/jgromero7/gos-twitter/jwt"
)

// CheckDataBase verify connectio with database
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// CheckJWT verify token client
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := jwtService.DecodeJWT(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
