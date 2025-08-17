package middleware

import (
	"backend/helper"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := retrieveRequestID(r)
		tokenString, err := request.BearerExtractor{}.ExtractToken(r)
		if err != nil {
			logrus.Error("Missing Authorization Header")
			helper.ResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayloadMissingToken)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			secretKey := []byte(os.Getenv("JWT_SECRET"))
			if len(secretKey) == 0 {
				return nil, errors.New("environment variable not found: jwt signing key")
			}
			return secretKey, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			logrus.Error(err)
			helper.ResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayloadInvalidToken)
			return
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if ok {
			logrus.Info("Token validated")
			next.ServeHTTP(w, r)
		} else {
			logrus.Info("Token validation failed")
			helper.ResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayloadInvalidToken)
			return
		}
	})
}

func retrieveRequestID(r *http.Request) (requestID uuid.UUID) {
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
