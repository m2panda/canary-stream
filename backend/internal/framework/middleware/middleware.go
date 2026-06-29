package middleware

import "net/http"

func MiddlewarePipeline(next http.Handler) http.Handler {
	return corsMiddleware(
		rateLimitMiddleware(next),
	)
}
