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

func BuildBadRequestPayload(requestUuid uuid.UUID, errorDetail *model.ErrorDetails) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "BAD_REQUEST",
		Message: "Invalid Request Format",
		Status:  http.StatusBadRequest,
		TraceID: requestUuid.String(),
		Details: errorDetail,
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func BuildInternalServerErrorPayload(requestUuid uuid.UUID, errorDetail *model.ErrorDetails) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Something went wrong. Please try again later.",
		Status:  http.StatusInternalServerError,
		TraceID: requestUuid.String(),
		Details: errorDetail,
	}})
	if err != nil {
		return responseBytes, err
	}
	return responseBytes, nil
}

func BuildUnauthorizedRequestPayload(requestUuid uuid.UUID, errorDetail *model.ErrorDetails) ([]byte, error) {
	responseBytes, err := json.Marshal(model.ErrorResponse{Error: model.APIError{
		Code:    "UNAUTHENTICATED",
		Message: "Authentication Token is missing, invalid or expired. Please log in again.",
		Status:  http.StatusUnauthorized,
		TraceID: requestUuid.String(),
		Details: errorDetail,
	}})
	if err != nil {
		return responseBytes, err
	}

	return responseBytes, nil
}

func ErrorResponseWriter(w http.ResponseWriter, statusCode int, requestID uuid.UUID, payloadBuilder func(uuid.UUID, *model.ErrorDetails) ([]byte, error), errorDetails *model.ErrorDetails) {
	SetResponseHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	responseBodyBytes, err := payloadBuilder(requestID, errorDetails)
	if err != nil {
		logrus.Error("Error creating response Body: ", err)
		return
	}
	_, err = w.Write(responseBodyBytes)
	if err != nil {
		logrus.Error("Error writing response: ", err)
	}
	return
}
