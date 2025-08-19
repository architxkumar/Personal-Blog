package middleware

import (
	"backend/helper"
	"backend/model"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/sirupsen/logrus"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := RetrieveRequestID(r)
		errorDetails := &model.ErrorDetails{Resource: "Dashboard"}
		tokenString, err := request.BearerExtractor{}.ExtractToken(r)
		if err != nil {
			logrus.Error("Missing Authorization Header")
			helper.ErrorResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayload, errorDetails)
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
			helper.ErrorResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayload, errorDetails)
			return
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if ok {
			logrus.Info("Token validated")
			next.ServeHTTP(w, r)
		} else {
			logrus.Info("Token validation failed")
			helper.ErrorResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayload, errorDetails)
			return
		}
	})
}
