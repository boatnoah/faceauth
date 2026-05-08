package main

import (
	"fmt"
	"net/http"
	"strings"
)

type userKey string

const userCtx userKey = "user"

func (a *app) AuthTokenMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header is missing", http.StatusBadRequest)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "authorization header is malformed", http.StatusBadRequest)
			return

		}

		token := parts[1]
		err := a.authenticator.ValidateToken(r.Context(), token)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

	})
}
