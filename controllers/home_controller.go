package controllers

import "net/http"

// HomeController struct
type HomeController struct {
	*Controller
}

// NewHomeController pointer of HomeController
func NewHomeController() *HomeController {
	return &HomeController{NewController()}
}

// Home page
func (c HomeController) Home(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("{\"message\": \"GeoLocation API\"}"))
}
