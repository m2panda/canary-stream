package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type ClientIP struct {
	limiter     *rate.Limiter
	lastRequest time.Time
}

var (
	mutex   sync.Mutex
	clients map[string]*ClientIP
)

func InitializeIPsCleaner() {
	clients = make(map[string]*ClientIP)

	go func() {
		for {
			time.Sleep(10 * time.Minute)
			mutex.Lock()

			for ip, client := range clients {
				if time.Since(client.lastRequest) > 15*time.Minute {
					delete(clients, ip)
				}
			}

			mutex.Unlock()
		}
	}()
}

func getLimiter(ip string) *rate.Limiter {
	mutex.Lock()
	defer mutex.Unlock()

	client, exists := clients[ip]

	if exists {
		client.lastRequest = time.Now()
		return client.limiter
	}

	limiter := rate.NewLimiter(rate.Every(time.Second/5), 5)

	clients[ip] = &ClientIP{
		limiter:     limiter,
		lastRequest: time.Now(),
	}

	return limiter
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		realIP := request.Header.Get("X-Forwarded-For")

		if realIP == "" {
			realIP, _, _ = net.SplitHostPort(request.RemoteAddr)
		}

		limiter := getLimiter(realIP)

		if !limiter.Allow() {
			http.Error(response, "Too many request", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(response, request)
	})
}
