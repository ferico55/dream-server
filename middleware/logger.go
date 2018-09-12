package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging provide common request logging
func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.Method, r.URL.Path, time.Since(start)) }()

			next(w, r)
		}
	}
}
