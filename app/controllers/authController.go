package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jgromero7/gos-twitter/app/models"
)

// SignUp for stoner one user
func SignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "The data sent is not correct", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "User email is required", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "The password must contain a minimum of 6 characters", http.StatusBadRequest)
		return
	}

	_, existUser, _ := models.ExistsUser(user.Email)
	if existUser == true {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	_, status, err := models.CreateUser(user)
	if err != nil {
		http.Error(w, "An error occurred while trying an error inserting a record", http.StatusConflict)
		return
	}

	if status == false {
		http.Error(w, "Cannot perform to user storage", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}
