package helper

import (
	"backend/model"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
