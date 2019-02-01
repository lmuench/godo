package api

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON parses v as JSON and responds with "OK" and v or "Bad Request"
func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	Handle500(w, err)
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// Handle400 responds with "Bad Request" status if err != nil
func Handle400(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Handle500 responds with "Internal Server Error" status if err != nil
func Handle500(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
