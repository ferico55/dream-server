package middleware

import "net/http"

// Middleware type definition
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain apply middlewares chaining to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
