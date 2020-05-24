package routes

import (
	"github.com/gorilla/mux"
	"github.com/jgromero7/gos-twitter/app/controllers"
	"github.com/jgromero7/gos-twitter/middleware"
)

// Routes config server
func Routes() *mux.Router {
	router := mux.NewRouter()

	// Users
	router.HandleFunc("/signUp", middleware.CheckDataBase(controllers.SignUp)).Methods("POST")
	router.HandleFunc("/signIn", middleware.CheckDataBase(controllers.SignIn)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckDataBase(middleware.CheckJWT(controllers.Profile))).Methods("GET")
	router.HandleFunc("/profile", middleware.CheckDataBase(middleware.CheckJWT(controllers.UpdateProfile))).Methods("PUT")

	// Upload Images for User
	router.HandleFunc("/profile/avatar", middleware.CheckDataBase(controllers.GetAvatar)).Methods("GET")
	router.HandleFunc("/profile/upload/avatar", middleware.CheckDataBase(middleware.CheckJWT(controllers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/profile/banner", middleware.CheckDataBase(controllers.GetBanner)).Methods("GET")
	router.HandleFunc("/profile/upload/banner", middleware.CheckDataBase(middleware.CheckJWT(controllers.UploadBanner))).Methods("POST")

	// Tweets
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.ReadTweet))).Methods("GET")
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.CreateTweet))).Methods("POST")
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.DeleteTweet))).Methods("DELETE")

	return router
}
