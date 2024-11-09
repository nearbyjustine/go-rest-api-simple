package main

import (
	"net/http"
	"strings"
)

func TokenAuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientId = r.URL.Query().Get("clientId")
		clientProfile, ok := database[clientId]
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token := r.Header.Get("Authorization")
		if !isValidToken(clientProfile, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}