package api

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON parses v as JSON and responds with OK and v
func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
