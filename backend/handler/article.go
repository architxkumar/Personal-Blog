package handler

import (
	"backend/helper"
	"backend/middleware"
	"backend/model"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ArticleCreationHandler(w http.ResponseWriter, r *http.Request) {
	requestID := middleware.RetrieveRequestID(r)
	errorDetails := &model.ErrorDetails{Resource: "Article"}
	articleDTO, ok := r.Context().Value("article").(model.ArticleWithoutIdDTO)
	if !ok {
		logrus.Error("Error retrieving ArticleDTO from context")
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, requestID, helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	articleUUID := uuid.New()
	newArticleEntry := model.Article{Id: articleUUID, Title: articleDTO.Title, Content: articleDTO.Content, Date: articleDTO.Date}
	var articles []model.Article
	fileContentBytes, err := os.ReadFile(helper.ArticleFilePath)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	logrus.Debug("Local file loaded successfully")
	err = json.Unmarshal(fileContentBytes, &articles)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	logrus.Debug("Contents read successfully")
	articles = append(articles, newArticleEntry)
	fileContentBytes, err = json.Marshal(articles)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	err = os.WriteFile(helper.ArticleFilePath, fileContentBytes, 0755)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	logrus.Info("Article created successfully")
	newArticleEntryBytes, err := json.Marshal(newArticleEntry)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
	helper.SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(newArticleEntryBytes)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, errorDetails)
		return
	}
}
