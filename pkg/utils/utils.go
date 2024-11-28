package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// TODO: Move this to HTTP pkg, rename to Decode and use Generics
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
