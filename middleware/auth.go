package middleware

import (
	"net/http"
)

// Auth provide auth checking before continuing
func Auth() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
			if true {
				next(w, r)
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusBadRequest)
				return
			}
		}
	}
}
