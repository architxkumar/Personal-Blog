package helper

import (
	"backend/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
}

func BuildBadRequestPayload(requestUuid uuid.UUID) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "BAD_REQUEST",
		Message: "Invalid Request Format",
		Status:  http.StatusBadRequest,
		TraceID: requestUuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func BuildInternalServerErrorPayload(requestUuid uuid.UUID) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal Server Error",
		Status:  http.StatusInternalServerError,
		TraceID: requestUuid.String(),
	}})
	if err != nil {
		return responseBytes, err
	}
	return responseBytes, nil
}

func BuildUnauthorizedRequestPayload(requestUuid uuid.UUID) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "INVALID_CREDENTIALS",
		Message: "Invalid email or password",
		Status:  http.StatusUnauthorized,
		TraceID: requestUuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func BuildUnauthorizedRequestPayloadMissingToken(requestUuid uuid.UUID) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "INVALID_CREDENTIALS",
		Message: "Missing Token",
		Status:  http.StatusUnauthorized,
		TraceID: requestUuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func BuildUnauthorizedRequestPayloadInvalidToken(requestUuid uuid.UUID) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "INVALID_CREDENTIALS",
		Message: "Invalid Token. Please login again",
		Status:  http.StatusUnauthorized,
		TraceID: requestUuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func ResponseWriter(w http.ResponseWriter, statusCode int, requestID uuid.UUID, payloadBuilder func(uuid2 uuid.UUID) ([]byte, error)) {
	w.WriteHeader(statusCode)
	responseBodyBytes, err := payloadBuilder(requestID)
	if err != nil {
		logrus.Error("Error creating response Body: ", err)
		return
	}
	SetResponseHeaders(w)
	_, err = w.Write(responseBodyBytes)
	if err != nil {
		logrus.Error("Error writing response: ", err)
	}
	return
}
