package routes

import (
	"github.com/gorilla/mux"
	"github.com/jgromero7/gos-twitter/app/controllers"
	"github.com/jgromero7/gos-twitter/middleware"
)

// Routes config server
func Routes() *mux.Router {
	router := mux.NewRouter()

	// Auth & Register
	router.HandleFunc("/signUp", middleware.CheckDataBase(controllers.SignUp)).Methods("POST")
	router.HandleFunc("/signIn", middleware.CheckDataBase(controllers.SignIn)).Methods("POST")

	// Users
	router.HandleFunc("/profile", middleware.CheckDataBase(middleware.CheckJWT(controllers.Profile))).Methods("GET")
	router.HandleFunc("/profile", middleware.CheckDataBase(middleware.CheckJWT(controllers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/profile/highUser", middleware.CheckDataBase(middleware.CheckJWT(controllers.HighUser))).Methods("GET")
	router.HandleFunc("/profile/lowUser", middleware.CheckDataBase(middleware.CheckJWT(controllers.LowUser))).Methods("DELETE")
	router.HandleFunc("/profile/queryRelation", middleware.CheckDataBase(middleware.CheckJWT(controllers.QueryRealtion))).Methods("GET")
	router.HandleFunc("/profile/usersList", middleware.CheckDataBase(middleware.CheckJWT(controllers.ListingUsers))).Methods("GET")

	// Upload Images for User
	router.HandleFunc("/profile/avatar", middleware.CheckDataBase(controllers.GetAvatar)).Methods("GET")
	router.HandleFunc("/profile/upload/avatar", middleware.CheckDataBase(middleware.CheckJWT(controllers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/profile/banner", middleware.CheckDataBase(controllers.GetBanner)).Methods("GET")
	router.HandleFunc("/profile/upload/banner", middleware.CheckDataBase(middleware.CheckJWT(controllers.UploadBanner))).Methods("POST")

	// Tweets
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.ReadTweet))).Methods("GET")
	router.HandleFunc("/tweet/followers", middleware.CheckDataBase(middleware.CheckJWT(controllers.ReadTweetFollowers))).Methods("GET")
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.CreateTweet))).Methods("POST")
	router.HandleFunc("/tweet", middleware.CheckDataBase(middleware.CheckJWT(controllers.DeleteTweet))).Methods("DELETE")

	return router
}
