package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// Middleware for JWT authentication
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT from the cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized: No token found", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		// Parse the token
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}

// NewCORSMiddleware adds CORS headers and handles OPTIONS preflight requests
func CORSMiddleware() func(http.Handler) http.Handler {
	log := zerolog.Ctx(context.Background())
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

			// Handle preflight request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
			log.Println("CORS middleware triggered for", r.Method, r.URL.Path)
		}

		return http.HandlerFunc(fn)
	}
}
