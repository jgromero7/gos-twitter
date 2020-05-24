package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jgromero7/gos-twitter/app/models"
	jwtServices "github.com/jgromero7/gos-twitter/jwt"
)

// CreateTweet stora a register tweet
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Content Invalid", http.StatusBadRequest)
		return
	}

	auxTweet := models.Tweet{
		UserID:   jwtServices.UserID,
		Message:  tweet.Message,
		CreateAt: time.Now(),
	}

	_, status, err := models.Create(auxTweet)
	if err != nil {
		http.Error(w, "An error occurred while trying an error inserting a record", http.StatusConflict)
		return
	}

	if status == false {
		http.Error(w, "Cannot perform to user storage", http.StatusConflict)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ReadTweet Listing all result
func ReadTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	UserID := r.URL.Query().Get("userid")

	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "The param page is required", http.StatusNotAcceptable)
		return
	}

	Page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "the page parameter must contain a value", http.StatusNotAcceptable)
		return
	}

	numberPage := int64(Page)
	tweets, status := models.Read(UserID, numberPage)
	if status == false {
		http.Error(w, "error reading tweets", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

// DeleteTweet delete a register tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The param id is required", http.StatusNotAcceptable)
		return
	}

	UserID := r.URL.Query().Get("userid")

	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	err := models.Delete(ID, UserID)
	if err != nil {
		http.Error(w, "error deleting tweet", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
