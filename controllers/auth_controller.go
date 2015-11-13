package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/strom87/ApiGeoTracking/core/factory"
	"github.com/strom87/ApiGeoTracking/models"
)

// AuthController struct
type AuthController struct {
	*Controller
}

// NewAuthController pointer of AuthController
func NewAuthController() *AuthController {
	return &AuthController{NewController()}
}

// Login login user
func (c AuthController) Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	model := &models.LoginModel{}
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		c.Logger.Log("AuthController login json decode.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 1})
		return
	}

	conn := factory.NewDatabase().Connection()
	if err := conn.Open(); err != nil {
		c.Logger.Log("AuthController login could not open db connection.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 2})
		return
	}
	defer conn.Close()

	user, err := factory.NewDatabase().User(conn).FindByEmail(model.Email)
	if err != nil || user == nil {
		c.Logger.Log("AuthController login find by email.", err)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 3})
		return
	}

	if user.Email != strings.ToLower(model.Email) || user.Password != model.Password {
		c.Logger.Log("AuthController login invalid username or password.", nil)
		json.NewEncoder(w).Encode(models.Response{ErrorCode: 4})
		return
	}

	json.NewEncoder(w).Encode(models.Response{Data: user, Success: true})
}
