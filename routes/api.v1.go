package routes

import (
	"services/jenkins_example/app/controller"
	"services/jenkins_example/app/middleware"

	"github.com/gorilla/mux"
)

/**
 * Definition of all routes in the api version 1
 * @return mux.Router
 */
func AplicationV1Router() *mux.Router {
	router := mux.NewRouter()
	// set routes
	router.HandleFunc("/book", controller.CreateBoook).Methods("GET")
	router.Use(middleware.LogginMiddleware)

	return router
}