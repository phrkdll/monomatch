package utils

import (
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
