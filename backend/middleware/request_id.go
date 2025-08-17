package middleware

import (
	"net/http"

	uuid2 "github.com/google/uuid"
	"golang.org/x/net/context"
)

type key string

const requestIDKey key = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid2.New()

		ctx := context.WithValue(r.Context(), requestIDKey, requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
