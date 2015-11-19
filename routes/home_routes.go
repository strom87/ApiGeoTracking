package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/controllers"
)

func homeRoutes(router *mux.Router) {
	controller := controllers.NewHomeController()

	router.Handle("/", negroni.New(
		negroni.HandlerFunc(controller.Home),
	)).Methods("GET")
}
