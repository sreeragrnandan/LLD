package main

import (
	"fmt"
	"net/http"
	"rate-limiter/ratelimiter"
	"time"
)

func main() {
	rateLimiter := ratelimiter.NewRateLimiter(10, 6*time.Second)

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userID")
		if userID == "" {
			http.Error(w, "missing user ID", http.StatusBadRequest)
			return
		}

		if rateLimiter.Allow(userID) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Request allowed"))
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rete Limit exceeded"))
		}
	})
	fmt.Println("Listening to server http://localhost:8080/check?userID=?")
	http.ListenAndServe(":8080", nil)
}
