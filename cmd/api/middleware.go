package main

import (
	"context"
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
		sessionToken, err := app.authenticator.ValidateToken(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		user, err := app.getUser(ctx, userID)
		if err != nil {
			http.Error(w, "Unable to retrieve user", http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *app) getUser(ctx context.Context, userID int64) (*storage.User, error) {

	// todo

	user, err := app.store.UserStorage.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *app) getUserFromContext(r *http.Request) *storage.User {
	user, _ := r.Context().Value(userCtx).(*storage.User)
	return user
}
