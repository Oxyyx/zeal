package render

import (
	"encoding/json"
	"net/http"
)

type JSONError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

// JSON writes any input interface{} (v) to the input http.ResponseWriter after Marshalling it to JSON.
// HTTP Status code will be set based on input "status" parameter.
// Content-Type will always be set to "application/json"
func JSON(w http.ResponseWriter, status int, v interface{}) {
	response, err := json.Marshal(v)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// JSONError writes a JSON Error to the input http.ResponseWriter based on status code and user provided messages.
// Example messages look like {"code": 400, "message": "Bad request"}
func JSONErr(w http.ResponseWriter, status int, message string) {
	JSON(w, status, JSONError{Code: status, Message: message})
}