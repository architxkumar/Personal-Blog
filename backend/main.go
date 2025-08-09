package main

import (
	"backend/handler"
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
	http.HandleFunc("/auth/admin/login", handler.AuthenticationHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatal("Error in ListenAndServe")
	}
	logrus.Info("Listening on port 8080")
}
