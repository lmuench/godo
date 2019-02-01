package middleware

import "net/http"

// ContentTypeJSON middleware adds "Content-Type: application/json" to header
func ContentTypeJSON(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Add("Content-Type", "application/json")
	next(w, r)
}

// CORS middleware sets "Access-Control-Allow-Origin *" in header
func CORS(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next(w, r)
}
