package auth

import (
	"context"
	"net/http"
	"strings"
)

// RequireAuth validates the Bearer token, injects *auth.User into ctx.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		user, err := authSvc.GetUser(r.Context(), parts[1])
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
