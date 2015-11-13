package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/controllers"
)

func authRoutes(router *mux.Router) {
	controller := controllers.NewAuthController()

	router.Handle("/login", negroni.New(
		negroni.HandlerFunc(controller.Login),
	)).Methods("POST")
}
