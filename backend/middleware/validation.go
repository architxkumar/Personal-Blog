package middleware

import (
	"backend/helper"
	"backend/model"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func Validation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := RetrieveRequestID(r)
		errorDetails := &model.ErrorDetails{Resource: "Article", Id: nil}
		contentTypeHeader := r.Header.Get("Content-Type")
		mimeType := strings.Split(contentTypeHeader, ";")[0]
		if strings.ToLower(mimeType) != "application/json" {
			logrus.Error("Validation Failed: Invalid content-type header ", mimeType)
			helper.ErrorResponseWriter(w, http.StatusUnsupportedMediaType, RetrieveRequestID(r), helper.BuildUnsupportedMediaTypeRequestPayload, errorDetails)
			return
		}
		var article model.ArticleWithoutIdDTO
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			logrus.WithField("Error", err).Error("Validation Failed: Error decoding request Body")
			helper.ErrorResponseWriter(w, http.StatusBadRequest, requestID, helper.BuildBadRequestPayload, errorDetails)
			return
		}

		if strings.TrimSpace(article.Title) == "" || strings.TrimSpace(article.Content) == "" || strings.TrimSpace(article.Date) == "" {
			logrus.WithFields(logrus.Fields{
				"Title":   article.Title,
				"Content": article.Content,
				"Date":    article.Date,
			}).Error("Validation Failed: Missing required fields")
			helper.ErrorResponseWriter(w, http.StatusBadRequest, requestID, helper.BuildBadRequestPayload, errorDetails)
			return
		}
		ctx := context.WithValue(r.Context(), "article", article)
		logrus.Info("Request Body Validated Successfully")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
