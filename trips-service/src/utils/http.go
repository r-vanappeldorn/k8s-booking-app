// Package utils:
package utils

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
