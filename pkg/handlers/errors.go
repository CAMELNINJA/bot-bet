package handlers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func AcseptError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Error{Message: err.Error()})
}
