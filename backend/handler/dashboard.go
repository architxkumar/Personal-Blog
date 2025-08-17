package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func DashBoardHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("I am inside the handler")
	w.WriteHeader(http.StatusCreated)
	return
}
