package middlewares

import "net/http"

// JSONHeader struct container
type JSONHeader struct{}

// NewJSONHeader get pointer to JsonHeader struct
func NewJSONHeader() *JSONHeader {
	return &JSONHeader{}
}

func (JSONHeader) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	next(w, req)
}
