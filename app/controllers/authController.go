package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jgromero7/gos-twitter/app/models"
	"github.com/jgromero7/gos-twitter/app/structs"
	jwtServices "github.com/jgromero7/gos-twitter/jwt"
)

// SignUp for store one user
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

// SignIn for auth user
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Content Invalid", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	currentUser, exists := models.CheckSignIn(user.Email, user.Password)
	if exists == false {
		http.Error(w, "Incorrect user or password", http.StatusBadRequest)
		return
	}

	payload := jwt.MapClaims{
		"_id":       currentUser.ID.Hex(),
		"name":      currentUser.Name,
		"lastName":  currentUser.LastName,
		"birthDate": currentUser.BirthDate,
		"email":     currentUser.Email,
		"avatar":    currentUser.Avatar,
		"banner":    currentUser.Banner,
		"biography": currentUser.Biography,
		"location":  currentUser.Location,
		"webSite":   currentUser.WebSite,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtKey, err := jwtServices.GenerateJWT(payload)
	if err != nil {
		http.Error(w, "the request could not be processed", http.StatusUnprocessableEntity)
		return
	}

	dataResponse := structs.ReponseSignIn{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataResponse)

	expirateTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirateTime,
	})
}

// Profile show info user
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The param id is required", http.StatusNotFound)
	}

	profile, err := models.GetUser(ID)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfile updated info user
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "The data sent is not correct", http.StatusBadRequest)
		return
	}

	// UserID proviene del JWT file decodeToken
	var status bool
	status, err = models.UpdateUser(user, jwtServices.UserID)
	if err != nil {
		http.Error(w, "An error occurred while updating the user", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Could not update user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
