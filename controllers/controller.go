package controllers

import "github.com/strom87/ApiGeoTracking/core/logger"

// Controller struct
type Controller struct {
	Logger *logger.Logger
}

// NewController pointer to Controller
func NewController() *Controller {
	return &Controller{Logger: logger.NewLogger()}
}
