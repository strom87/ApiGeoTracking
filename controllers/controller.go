package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/strom87/ApiGeoTracking/core/logger"
)

// Controller struct
type Controller struct {
	Logger *logger.Logger
}

// NewController pointer to Controller
func NewController() *Controller {
	return &Controller{Logger: logger.NewLogger()}
}

// GetString get int from url paramater by id
func (Controller) GetString(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}
