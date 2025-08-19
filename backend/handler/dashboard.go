package handler

import (
	"backend/helper"
	"backend/middleware"
	"backend/model"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func DashBoardHandler(w http.ResponseWriter, r *http.Request) {
	var articles []model.Article
	fileContents, err := os.ReadFile("data/articles.json")
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, nil)
		return
	}
	logrus.Debug("Local file loaded successfully")
	err = json.Unmarshal(fileContents, &articles)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, nil)
		return
	}
	logrus.Debug("Contents read successfully")
	var articlePreviewsDTO []model.ArticlePreviewDTO
	for _, article := range articles {
		articlePreviewsDTO = append(articlePreviewsDTO, model.ArticlePreviewDTO{Id: article.Id, Title: article.Title})
	}
	responseBytes, err := json.Marshal(articlePreviewsDTO)
	if err != nil {
		logrus.Error(err)
		helper.ErrorResponseWriter(w, http.StatusInternalServerError, middleware.RetrieveRequestID(r), helper.BuildInternalServerErrorPayload, nil)
		return
	}
	helper.SetResponseHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBytes)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("Response sent successfully")
}
