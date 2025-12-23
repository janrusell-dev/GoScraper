package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "The server is busy scraping. Please try again later.", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
