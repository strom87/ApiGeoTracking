package middlewares

import "net/http"

// OriginHeader struct container
type OriginHeader struct{}

// NewOriginHeader get pointer to JsonHeader struct
func NewOriginHeader() *OriginHeader {
	return &OriginHeader{}
}

func (OriginHeader) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next(w, req)
}
