package gorilla

import (
	"context"
	"net/http"
	"os"
)

const (
	login = "jeffreyz"
	pass  = "confuzed"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != login || password != pass {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := os.Getuid()

		ctx := context.WithValue(context.Background(), "REQUEST_ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
