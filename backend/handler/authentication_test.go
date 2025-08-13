package handler

import (
	"backend/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationHandler_FailedPath(t *testing.T) {
	t.Run("Invalid Method Error", func(t *testing.T) {
		loginRequestModel := model.LoginRequest{Email: "admin@example.com",
			Password: "admin",
		}
		requestBodyBytes, err := json.Marshal(&loginRequestModel)
		require.NoError(t, err, "Request Body Json Marshalling should return nil")
		req := httptest.NewRequest(http.MethodGet, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(AuthenticationHandler)
		handler.ServeHTTP(w, req)
		require.Equal(t, http.StatusMethodNotAllowed, w.Code)
		expectedHeaders := map[string]string{
			"Content-Type":            "text/plain",
			"X-Content-Type-Options":  "nosniff",
			"X-Frame-Options":         "deny",
			"Content-Security-Policy": "default-src 'none'",
		}
		for key, value := range expectedHeaders {
			assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
		}
		assert.Empty(t, w.Body.String(), "Response Body must be empty")
	})
	t.Run("Bad Status Request Error", func(t *testing.T) {
		t.Run("Missing Fields", func(t *testing.T) {
			loginRequestModel := model.LoginRequest{}
			requestBodyBytes, err := json.Marshal(&loginRequestModel)
			require.NoError(t, err, "Request Body Json Marshalling should return nil")
			req := httptest.NewRequest(http.MethodPost, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(AuthenticationHandler)
			handler.ServeHTTP(w, req)
			require.Equal(t, http.StatusBadRequest, w.Code)
			expectedHeaders := map[string]string{
				"Content-Type":            "application/json",
				"X-Content-Type-Options":  "nosniff",
				"X-Frame-Options":         "deny",
				"Content-Security-Policy": "default-src 'none'",
			}
			for key, value := range expectedHeaders {
				assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
			}
			var responseBody model.ErrorResponse
			err = json.Unmarshal(w.Body.Bytes(), &responseBody)
			require.NoError(t, err, "Response Body unmarshalling should return nil")
			assert.Equal(t, "BAD_REQUEST", responseBody.Error.Code)
			assert.Equal(t, "Invalid Request Format", responseBody.Error.Message)
			assert.Equal(t, http.StatusBadRequest, responseBody.Error.Status)
			assert.NotEmpty(t, responseBody.Error.TraceID)
			assert.Equal(t, "Login", responseBody.Error.Details.Resource)
			assert.Nil(t, responseBody.Error.Details.Id)
		})
		t.Run("Missing Email", func(t *testing.T) {
			loginRequestModel := model.LoginRequest{Password: "admin"}
			requestBodyBytes, err := json.Marshal(&loginRequestModel)
			require.NoError(t, err, "Request Body Json Marshalling should return nil")
			req := httptest.NewRequest(http.MethodPost, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(AuthenticationHandler)
			handler.ServeHTTP(w, req)
			require.Equal(t, http.StatusBadRequest, w.Code)
			expectedHeaders := map[string]string{
				"Content-Type":            "application/json",
				"X-Content-Type-Options":  "nosniff",
				"X-Frame-Options":         "deny",
				"Content-Security-Policy": "default-src 'none'",
			}
			for key, value := range expectedHeaders {
				assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
			}
			var responseBody model.ErrorResponse
			err = json.Unmarshal(w.Body.Bytes(), &responseBody)
			require.NoError(t, err, "Response Body unmarshalling should return nil")
			assert.Equal(t, "BAD_REQUEST", responseBody.Error.Code)
			assert.Equal(t, "Invalid Request Format", responseBody.Error.Message)
			assert.Equal(t, http.StatusBadRequest, responseBody.Error.Status)
			assert.NotEmpty(t, responseBody.Error.TraceID)
			assert.Equal(t, "Login", responseBody.Error.Details.Resource)
			assert.Nil(t, responseBody.Error.Details.Id)
		})
		t.Run("Missing Password", func(t *testing.T) {
			loginRequestModel := model.LoginRequest{Email: "admin@example.com"}
			requestBodyBytes, err := json.Marshal(&loginRequestModel)
			require.NoError(t, err, "Request Body Json Marshalling should return nil")
			req := httptest.NewRequest(http.MethodPost, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(AuthenticationHandler)
			handler.ServeHTTP(w, req)
			require.Equal(t, http.StatusBadRequest, w.Code)
			expectedHeaders := map[string]string{
				"Content-Type":            "application/json",
				"X-Content-Type-Options":  "nosniff",
				"X-Frame-Options":         "deny",
				"Content-Security-Policy": "default-src 'none'",
			}
			for key, value := range expectedHeaders {
				assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
			}
			var responseBody model.ErrorResponse
			err = json.Unmarshal(w.Body.Bytes(), &responseBody)
			require.NoError(t, err, "Response Body unmarshalling should return nil")
			assert.Equal(t, "BAD_REQUEST", responseBody.Error.Code)
			assert.Equal(t, "Invalid Request Format", responseBody.Error.Message)
			assert.Equal(t, http.StatusBadRequest, responseBody.Error.Status)
			assert.NotEmpty(t, responseBody.Error.TraceID)
			assert.Equal(t, "Login", responseBody.Error.Details.Resource)
			assert.Nil(t, responseBody.Error.Details.Id)
		})
	})
	t.Run("Internal Server Error", func(t *testing.T) {
		loginRequestModel := model.LoginRequest{Email: "admin@example.com", Password: "admin"}
		requestBodyBytes, err := json.Marshal(&loginRequestModel)
		require.NoError(t, err, "Request Body Json Marshalling should return nil")
		req := httptest.NewRequest(http.MethodPost, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(AuthenticationHandler)
		handler.ServeHTTP(w, req)
		require.Equal(t, http.StatusInternalServerError, w.Code)
		expectedHeaders := map[string]string{
			"Content-Type":            "application/json",
			"X-Content-Type-Options":  "nosniff",
			"X-Frame-Options":         "deny",
			"Content-Security-Policy": "default-src 'none'",
		}
		for key, value := range expectedHeaders {
			assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
		}
		var responseBody model.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &responseBody)
		require.NoError(t, err, "Response Body unmarshalling should return nil")
		assert.Equal(t, "INTERNAL_SERVER_ERROR", responseBody.Error.Code)
		assert.Equal(t, "Internal Server Error", responseBody.Error.Message)
		assert.Equal(t, http.StatusInternalServerError, responseBody.Error.Status)
		assert.NotEmpty(t, responseBody.Error.TraceID)
		assert.Nil(t, responseBody.Error.Details)
	})
}

func TestAuthenticationHandler_SuccessPath(t *testing.T) {

	loginRequestModel := model.LoginRequest{Email: "admin@example.com",
		Password: "admin",
	}
	err := godotenv.Load("../.env")
	require.NoError(t, err, "Env file loading should return nil")
	requestBodyBytes, err := json.Marshal(&loginRequestModel)
	require.NoError(t, err, "Request Body Json Marshalling should return nil")
	req := httptest.NewRequest(http.MethodPost, "/auth/admin/login", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(AuthenticationHandler)

	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	expectedHeaders := map[string]string{
		"Content-Type":            "application/json",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "deny",
		"Content-Security-Policy": "default-src 'none'",
	}
	for key, value := range expectedHeaders {
		assert.Equal(t, value, w.Header().Get(key), fmt.Sprintf("%s header", key))
	}
	var responsePayload model.LoginResponseDTO
	err = json.Unmarshal(w.Body.Bytes(), &responsePayload)
	require.NoError(t, err, "Response Body Json Unmarshalling should return nil")
	assert.NotEmpty(t, responsePayload.Token, "Response Payload should not be nil")
}
