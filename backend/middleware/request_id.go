package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type key string

const requestIDKey key = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New()

		ctx := context.WithValue(r.Context(), requestIDKey, requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RetrieveRequestID(r *http.Request) (requestID uuid.UUID) {
	value := r.Context().Value(requestIDKey)
	if value != nil {
		reqID, ok := value.(uuid.UUID)
		if !ok {
			logrus.Warn("Type assertion failure on 'requestID'")
		} else {
			requestID = reqID
		}
	} else {
		logrus.Warn("Missing requestID key")
	}
	return
}
