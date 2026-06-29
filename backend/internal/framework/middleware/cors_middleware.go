package middleware

import (
	"log/slog"
	"net/http"
	"strings"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		origin := request.Header.Get("Origin")

		if origin == "" {
			slog.Info("Unknown request origin",
				"event", "cors-origin",
			)

			next.ServeHTTP(response, request)
			return
		}

		host := request.Host
		origin = strings.TrimPrefix(origin, "http://")
		origin = strings.TrimPrefix(origin, "https://")

		if host != origin {
			slog.Error("Request origin not allowed",
				"event", "cors-origin",
				"origin", origin,
			)

			http.Error(response, "CORS: Origin not allowed (Same-Origin policy active)", http.StatusForbidden)
			return
		}

		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		response.Header().Set("Access-Control-Allow-Credentials", "true")

		if request.Method == http.MethodOptions {
			slog.Info("Preflight request",
				"event", "cors-allowed",
				"method", request.Method,
			)

			response.WriteHeader(http.StatusOK)
			return
		}

		slog.Info("CORS access verified",
			"event", "cors-allowed",
			"method", request.Method,
		)

		next.ServeHTTP(response, request)
	})
}
