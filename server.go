package main

import (
	"log"
	"net/http"

	"github.com/strom87/ApiGeoTracking/routes"
)

const port = ":1337"

func main() {
	log.Println("Listening on port", port)

	negroni := routes.InitRoutes()

	if err := http.ListenAndServe(port, negroni); err != nil {
		log.Fatal(err)
	}
}
