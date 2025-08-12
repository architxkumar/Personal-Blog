package handler

import (
	"backend/helper"
	"backend/model"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func AuthenticationHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	requestUuid := uuid.New()
	logrus.WithFields(logrus.Fields{
		"request_uuid":   requestUuid.String(),
		"method":         req.Method,
		"url":            req.URL.String(),
		"remote_addr":    req.RemoteAddr,
		"user_agent":     req.UserAgent(),
		"content_length": req.ContentLength,
	}).Info("Incoming request")
	if req.Method != http.MethodPost {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("Content-Security-Policy", "default-src 'none'")
		w.WriteHeader(http.StatusMethodNotAllowed)
		logrus.Error("Invalid method")
		logrus.WithFields(logrus.Fields{
			"request_uuid":  requestUuid.String(),
			"status_code":   http.StatusMethodNotAllowed,
			"response_time": time.Since(start),
			"response_size": 0,
		}).Info("Outgoing response")
		return
	}
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		logrus.Error("Invalid content type: ", contentType)
		helper.SetResponseHeaders(w)
		w.WriteHeader(http.StatusBadRequest)
		responseBytes, err := helper.BuildBadRequestPayload(requestUuid)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")
			return
		}
		_, err = w.Write(responseBytes)
		if err != nil {
			logrus.WithError(err).Error("Failed to write response")
		} else {
			logrus.WithFields(logrus.Fields{
				"request_uuid":  requestUuid.String(),
				"status_code":   http.StatusBadRequest,
				"response_time": time.Since(start),
				"response_size": len(responseBytes),
			}).Info("Outgoing response")
		}
		return
	}
	reqBody := req.Body
	var loginRequestDTO model.LoginRequest
	err := json.NewDecoder(reqBody).Decode(&loginRequestDTO)
	if err != nil {
		logrus.WithError(err).Error("Error decoding body")
		helper.SetResponseHeaders(w)
		w.WriteHeader(http.StatusBadRequest)
		responseBytes, err := helper.BuildBadRequestPayload(requestUuid)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")
			return
		}
		_, err = w.Write(responseBytes)
		if err != nil {
			logrus.WithError(err).Error("Failed to write response")
		} else {
			logrus.WithFields(logrus.Fields{
				"request_uuid":  requestUuid.String(),
				"status_code":   http.StatusBadRequest,
				"response_time": time.Since(start),
				"response_size": len(responseBytes),
			}).Info("Outgoing response")
		}
		return
	}
	err = validator.New().Struct(loginRequestDTO)
	if err != nil {
		logrus.WithError(err).Error("Error validating request")
		helper.SetResponseHeaders(w)
		w.WriteHeader(http.StatusBadRequest)
		responseBytes, err := helper.BuildBadRequestPayload(requestUuid)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")
			return
		}
		_, err = w.Write(responseBytes)
		if err != nil {
			logrus.WithError(err).Error("Failed to write response")
		} else {
			logrus.WithFields(logrus.Fields{
				"request_uuid":  requestUuid.String(),
				"status_code":   http.StatusBadRequest,
				"response_time": time.Since(start),
				"response_size": len(responseBytes),
			}).Info("Outgoing response")
		}
		return
	}
	if loginRequestDTO.Email != "admin@example.com" || loginRequestDTO.Password != "admin" {
		logrus.Error("Invalid Credentials")
		helper.SetResponseHeaders(w)
		w.WriteHeader(http.StatusUnauthorized)
		responseBytes, err := helper.BuildUnauthorizedRequestPayload(requestUuid)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshall response")
			return
		}
		_, err = w.Write(responseBytes)
		if err != nil {
			logrus.WithError(err).Error("Failed to write response")
		} else {
			logrus.WithFields(logrus.Fields{
				"request_uuid":  requestUuid.String(),
				"status_code":   http.StatusUnauthorized,
				"response_time": time.Since(start),
				"response_size": len(responseBytes),
			}).Info("Outgoing response")
		}
		return
	}
	key := []byte(os.Getenv("JWT_SECRET"))
	currentTime := time.Now().Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"iat": currentTime, "exp": currentTime + 3600})
	s, err := t.SignedString(key)
	if err != nil || len(key) == 0 {
		logrus.WithError(err).Error("Error signing token: ")
		helper.SetResponseHeaders(w)
		w.WriteHeader(http.StatusInternalServerError)
		responseBytes, err := helper.BuildInternalServerErrorPayload(requestUuid)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal response")
			return
		}
		_, err = w.Write(responseBytes)
		if err != nil {
			logrus.WithError(err).Error("Failed to write response")
		} else {
			logrus.WithFields(logrus.Fields{
				"request_uuid":  requestUuid.String(),
				"status_code":   http.StatusInternalServerError,
				"response_time": time.Since(start),
				"response_size": len(responseBytes),
			}).Info("Outgoing response")
		}
		return
	}
	logrus.Debugf("JWT Signing Token: %s", s)
	helper.SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	responseBytes, err := json.Marshal(model.LoginResponseDTO{Token: s})
	_, err = w.Write(responseBytes)
	if err != nil {
		logrus.WithError(err).Error("Error writing response")
	} else {
		logrus.Info("Authentication Success")
		logrus.WithFields(logrus.Fields{
			"request_uuid":  requestUuid.String(),
			"status_code":   http.StatusOK,
			"response_time": time.Since(start),
			"response_size": len(responseBytes),
		}).Info("Outgoing response")
	}
	return
}
