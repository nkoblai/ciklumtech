package common

import (
	"encoding/json"
	"net/http"
)

func ResponseWithStatusAndMessage(w http.ResponseWriter, status int, msg []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(msg)
}

func ResponseWithError(w http.ResponseWriter, httpError error) {
	errMsg, err := json.Marshal(httpError)
	if err != nil {
		ResponseWithStatusAndMessage(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	ResponseWithStatusAndMessage(w, http.StatusInternalServerError, errMsg)
}
