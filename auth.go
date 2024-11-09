package main

import (
	c "context"
	"net/http"
	"strings"
)

type key string

const CLIENT_PROFILE_KEY key = "clientProfile"

func TokenAuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientProfile := r.Context().Value("clientProfile").(ClientProfile)

		token := r.Header.Get("Authorization")
		if !isValidToken(clientProfile, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		ctx := c.WithValue(r.Context(), CLIENT_PROFILE_KEY, clientProfile)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}
