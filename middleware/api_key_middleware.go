package middleware

import (
	"net/http"

	"110yards.ca/libs/go/core/api"
	"110yards.ca/libs/go/core/logger"
)

var configuredKey string

func CreateApiKeyMiddleware(withKey string) func(http.Handler) http.Handler {
	// withKey must not be empty
	if withKey == "" {
		logger.Fatal("Configured API Key must not be empty")
	}

	configuredKey = withKey
	return apiKeyMiddleware
}

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		anonymousRoutes := map[string]bool{
			"/health": true,
		}

		// anonymous routes
		if anonymousRoutes[r.URL.Path] {
			logger.Infof("%s route allows anonymous access", r.URL.Path)
			next.ServeHTTP(w, r)
			return
		}

		key := r.Header.Get("Authorization")

		if key != configuredKey {
			logger.Errorf("Invalid API key: %s", key)
			api.Unauthorized(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
