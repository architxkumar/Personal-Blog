package main

import (
	"backend/handler"
	"backend/middleware"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	logrus.Debug(".env file loaded")
	http.HandleFunc(" /auth/admin/login", handler.AuthenticationHandler)
	http.Handle("GET /admin/dashboard", middleware.RequestIDMiddleware(middleware.Logging(middleware.Authentication(http.HandlerFunc(handler.DashBoardHandler)))))
	http.Handle("POST /articles", middleware.RequestIDMiddleware(middleware.Logging(middleware.Authentication(middleware.Validation(http.HandlerFunc(handler.ArticleCreationHandler))))))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatal("Error in ListenAndServe")
	}
	logrus.Info("Listening on port 8080")
}
