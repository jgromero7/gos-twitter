package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jgromero7/gos-twitter/app/models"
	"github.com/jgromero7/gos-twitter/app/structs"
	jwtServices "github.com/jgromero7/gos-twitter/jwt"
)

// Profile show info user
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The param id is required", http.StatusNotFound)
		return
	}

	profile, err := models.GetUser(ID)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
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

// HighUser from High User
func HighUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	UserID := r.URL.Query().Get("userid")
	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	var relation models.Relation
	relation.UserID = jwtServices.UserID
	relation.UserRelation = UserID

	status, err := models.CreateRelation(relation)
	if err != nil {
		http.Error(w, "an error occurred in store relation", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "an error occurred creating relation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// LowUser from High User
func LowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	UserID := r.URL.Query().Get("userid")
	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	var relation models.Relation
	relation.UserID = jwtServices.UserID
	relation.UserRelation = UserID

	status, err := models.DeleteRelation(relation)
	if err != nil {
		http.Error(w, "an error occurred in deleting relation", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "an error occurred deleting a relation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// QueryRealtion get relation from determinate user
func QueryRealtion(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	UserID := r.URL.Query().Get("userid")
	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	var relation models.Relation
	relation.UserID = jwtServices.UserID
	relation.UserRelation = UserID

	var response structs.ResponseQueryRelation
	status, err := models.ReadRelation(relation)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ListingUsers get all users
func ListingUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	typeUser := r.URL.Query().Get("type")
	search := r.URL.Query().Get("search")

	Page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "the page parameter must contain a value", http.StatusNotAcceptable)
		return
	}

	numberPage := int64(Page)

	users, status := models.ReadAllUsers(jwtServices.UserID, numberPage, search, typeUser)
	if status == false {
		http.Error(w, "an error occurred while reading the users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
