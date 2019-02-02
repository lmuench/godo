package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// RespondWithJSON parses v as JSON and responds with "OK" and v or "Bad Request"
func RespondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		logWithTimestamp(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func logWithTimestamp(err error) {
	log.Println(time.Now().Format(time.RFC850), err.Error())
}
