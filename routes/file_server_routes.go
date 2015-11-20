package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/controllers"
)

func fileServeRoutes(router *mux.Router) {
	controller := controllers.NewFileServeController()

	router.Handle("/image/{folder}/{file}", negroni.New(
		negroni.HandlerFunc(controller.GetImage),
	)).Methods("GET")
}
