package middlewares

import (
	"errors"
	"net/http"

	"github.com/Muhammad-Tounsi/VoteGo/api/auth"
	"github.com/Muhammad-Tounsi/VoteGo/api/responses"
)

// SetMiddlewareJSON : This will format all responses to JSON
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication : will check for the validity of the authentication token provided
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
