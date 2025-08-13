package helper

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid2 "github.com/google/uuid"
	assert2 "github.com/stretchr/testify/assert"
)

func TestSetResponseHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	SetResponseHeaders(w)
	assert := assert2.New(t)
	response := w.Result()

	expectedHeaders := map[string]string{
		"Content-Type":            "application/json",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "deny",
		"Content-Security-Policy": "default-src 'none'",
	}
	for key, value := range expectedHeaders {
		assert.Equal(value, response.Header.Get(key), fmt.Sprintf("%s header", key))
	}
	assert.Len(response.Header, len(expectedHeaders), "Unexpected header set")
}

func TestBuildBadRequestPayload(t *testing.T) {
	assert := assert2.New(t)
	uuid, err := uuid2.NewUUID()
	assert.NoError(err, "UUID Generation should return nil")
	responseBytes, err := BuildBadRequestPayload(uuid)
	assert.NoError(err, "Bad Request Payload Builder should return nil")

	expectedPayload := model.ErrorResponse{Error: model.APIError{
		Code:    "BAD_REQUEST",
		Message: "Invalid Request Format",
		Status:  http.StatusBadRequest,
		TraceID: uuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}}
	var actualPayload model.ErrorResponse
	err = json.Unmarshal(responseBytes, &actualPayload)
	assert.NoError(err, "Json Unmarshalling should return nil")

	assert.Equal(expectedPayload, actualPayload)
}

func TestBuildInternalServerErrorPayload(t *testing.T) {
	assert := assert2.New(t)
	uuid, err := uuid2.NewUUID()
	assert.NoError(err, "UUID Generation should return nil")
	responseBytes, err := BuildInternalServerErrorPayload(uuid)
	assert.NoError(err, "Bad Request Payload Builder should return nil")

	expectedPayload := model.ErrorResponse{Error: model.APIError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal Server Error",
		Status:  http.StatusInternalServerError,
		TraceID: uuid.String(),
	}}
	var actualPayload model.ErrorResponse
	err = json.Unmarshal(responseBytes, &actualPayload)
	assert.NoError(err, "Json Unmarshalling should return nil")

	assert.Equal(expectedPayload, actualPayload)
}

func TestBuildUnauthorizedRequestPayload(t *testing.T) {
	assert := assert2.New(t)
	uuid, err := uuid2.NewUUID()
	assert.NoError(err, "UUID Generation should return nil")
	responseBytes, err := BuildUnauthorizedRequestPayload(uuid)
	assert.NoError(err, "Bad Request Payload Builder should return nil")

	expectedPayload := model.ErrorResponse{Error: model.APIError{
		Code:    "INVALID_CREDENTIALS",
		Message: "Invalid email or password",
		Status:  http.StatusUnauthorized,
		TraceID: uuid.String(),
		Details: &model.ErrorDetails{
			Resource: "Login",
		},
	}}
	var actualPayload model.ErrorResponse
	err = json.Unmarshal(responseBytes, &actualPayload)
	assert.NoError(err, "Json Unmarshalling should return nil")

	assert.Equal(expectedPayload, actualPayload)
}
