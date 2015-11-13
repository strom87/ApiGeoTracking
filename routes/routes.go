package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/middlewares"
)

func setRoutes(router *mux.Router) {
	authRoutes(router)
}

// InitRoutes initalizes the router
func InitRoutes() *negroni.Negroni {
	n := negroni.Classic()
	router := mux.NewRouter()
	n.Use(middlewares.NewJSONHeader())

	setRoutes(router)

	n.UseHandler(router)
	return n
}
