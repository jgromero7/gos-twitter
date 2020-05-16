package routes

import (
	"github.com/gorilla/mux"
	"github.com/jgromero7/gos-twitter/app/controllers"
	"github.com/jgromero7/gos-twitter/middleware"
)

// Routes config server
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/signUp", middleware.CheckDataBase(controllers.SignUp)).Methods("POST")
	router.HandleFunc("/signIn", middleware.CheckDataBase(controllers.SignIn)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckDataBase(middleware.CheckJWT(controllers.Profile))).Methods("GET")

	return router
}
