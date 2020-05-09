package middleware

import (
	"net/http"

	"github.com/jgromero7/gos-twitter/database"
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
