package middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(requestIDKey)
		if requestID == nil {
			logrus.Warn("Missing request ID Key")
		}
		requestID, ok := requestID.(uuid.UUID)
		if !ok {
			logrus.Warn("Type assertion failure on request ID Key")
		}
		logrus.WithFields(logrus.Fields{
			"request_uuid":   requestID,
			"method":         r.Method,
			"url":            r.URL.String(),
			"remote_addr":    r.RemoteAddr,
			"user_agent":     r.UserAgent(),
			"content_length": r.ContentLength,
		}).Info("Incoming request")
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.WithFields(logrus.Fields{
			"request_uuid":  requestID,
			"response_time": time.Since(start),
			"response_size": w.Header().Get("Content-Length"),
		}).Info("Outgoing response")

	})
}
