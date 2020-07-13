package web

import (
	"encoding/json"
	"net/http"
)

// ResponseWithStatusAndMessage writes to ResponseWriter with specific status code and message
func ResponseWithStatusAndMessage(w http.ResponseWriter, status int, msg []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(msg)
}

// ResponseWithError writes to ResponseWriter with specific error
func ResponseWithError(w http.ResponseWriter, httpError error) {
	errMsg, err := json.Marshal(httpError)
	if err != nil {
		ResponseWithStatusAndMessage(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	ResponseWithStatusAndMessage(w, http.StatusInternalServerError, errMsg)
}
