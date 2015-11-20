package controllers

import (
	"net/http"
	"os"
)

// FileServeController struct
type FileServeController struct {
	*Controller
}

// NewFileServeController returns a pointer of FileServeController
func NewFileServeController() *FileServeController {
	return &FileServeController{NewController()}
}

// GetImage writes an image to the http request
func (c FileServeController) GetImage(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	path := "./public/images/" + c.GetString(r, "folder") + "/" + c.GetString(r, "file")
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

	http.NotFound(w, r)
}
