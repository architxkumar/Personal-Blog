package middleware

import (
	"backend/helper"
	"backend/model"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/sirupsen/logrus"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := RetrieveRequestID(r)
		errorDetails := &model.ErrorDetails{Resource: "AdminDashboard"}
		tokenString, err := request.BearerExtractor{}.ExtractToken(r)
		if err != nil {
			logrus.Error("Missing Authorization Header")
			w.Header().Set("WWW-Authenticate", "Bearer")
			helper.ErrorResponseWriter(w, http.StatusUnauthorized, requestID, helper.BuildUnauthorizedRequestPayload, errorDetails)
			return
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		err = ValidateJWTSecretKey()
		if err != nil {
			logrus.Error(err)
			helper.ErrorResponseWriter(w, http.StatusInternalServerError, requestID, helper.BuildInternalServerErrorPayload, nil)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			jwtSecretBytes := []byte(jwtSecret)
			if len(jwtSecretBytes) == 0 {
				return nil, errors.New("environment variable not found: jwt signing key")
			}
			return jwtSecretBytes, nil
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

func ValidateJWTSecretKey() error {
	value, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return errors.New("missing jwt secret key variable")
	}
	value = strings.Trim(value, " ")
	if value == "" {
		return errors.New("jwt secrete key in empty")
	}
	return nil
}
