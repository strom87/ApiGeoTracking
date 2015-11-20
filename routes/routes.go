package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/middlewares"
	"github.com/strom87/ApiGeoTracking/websockets"
)

func setRoutes(router *mux.Router) {
	homeRoutes(router)
	authRoutes(router)
	fileServeRoutes(router)
}

func setWebsoketRoutes(router *mux.Router) {
	router.Handle("/ws/geo-location/{id}", negroni.New(
		websockets.NewLocationSocket(),
	))
}

// InitRoutes initalizes the router
func InitRoutes() *negroni.Negroni {
	n := negroni.Classic()
	router := mux.NewRouter()
	n.Use(middlewares.NewOriginHeader())
	n.Use(middlewares.NewJSONHeader())

	setRoutes(router)
	setWebsoketRoutes(router)

	n.UseHandler(router)
	return n
}
