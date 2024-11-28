package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Encode encodes the response and writes it
func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v *T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Encode json: %w", err)
	}

	return nil
}

