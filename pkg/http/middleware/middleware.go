package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// NewLoggingMiddleware logs the incoming HTTP request info
func NewLoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := zerolog.Ctx(context.Background())

			start := time.Now()

			// Call the next handler
			next.ServeHTTP(w, r)

			// post request stats
			log.Info().
				Str("method", r.Method).
				Str("path", r.URL.EscapedPath()).
				Int64("durationMS", time.Since(start).Milliseconds()).
				Msg("Request handled")
		}

		return http.HandlerFunc(fn)
	}
}

// NewPanicMiddleware recovers from panics and returns a server error
func NewPanicMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log := zerolog.Ctx(r.Context())
					log.
						Error().
						Interface("error", err).
						Msg("panic recovered")
				}
			}()

			// Call the next handler
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
