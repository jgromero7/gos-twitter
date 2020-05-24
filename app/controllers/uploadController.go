package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jgromero7/gos-twitter/app/models"
	jwtServices "github.com/jgromero7/gos-twitter/jwt"
)

// GetAvatar get img avatar from user
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	UserID := r.URL.Query().Get("userid")

	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	profile, err := models.GetUser(UserID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	OpenFile, err := os.Open("public/upload/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Avatars not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "error showing avatar", http.StatusBadRequest)
		return
	}
}

// UploadAvatar upload file avatar from user
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "error reading image"+err.Error(), http.StatusBadRequest)
		return
	}

	ext := strings.Split(handler.Filename, ".")[1]
	pathFile := "public/upload/avatars/" + jwtServices.UserID + "." + ext

	currentFile, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error loading image", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(currentFile, file)
	if err != nil {
		http.Error(w, "error loading image", http.StatusBadRequest)
		return
	}

	var user models.User
	user.Avatar = jwtServices.UserID + "." + ext
	var status bool
	status, err = models.UpdateUser(user, jwtServices.UserID)
	if err != nil || status == false {
		http.Error(w, "error storing avatar", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetBanner get img avatar from user
func GetBanner(w http.ResponseWriter, r *http.Request) {
	UserID := r.URL.Query().Get("userid")

	if len(UserID) < 1 {
		http.Error(w, "The param userid is required", http.StatusNotAcceptable)
		return
	}

	profile, err := models.GetUser(UserID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	OpenFile, err := os.Open("public/upload/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "error showing Banner", http.StatusBadRequest)
		return
	}
}

// UploadBanner upload file banner from user
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "error reading image"+err.Error(), http.StatusBadRequest)
		return
	}

	ext := strings.Split(handler.Filename, ".")[1]
	pathFile := "public/upload/banners/" + jwtServices.UserID + "." + ext

	currentFile, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error loading image", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(currentFile, file)
	if err != nil {
		http.Error(w, "error loading image", http.StatusBadRequest)
		return
	}

	var user models.User
	user.Banner = jwtServices.UserID + "." + ext
	var status bool
	status, err = models.UpdateUser(user, jwtServices.UserID)
	if err != nil || status == false {
		http.Error(w, "error storing banner", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
