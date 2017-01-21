package main

import (
	"net/http"

	"github.com/rumyantseva/sunday-go-school/02-web-server/handlers"
	"github.com/Sirupsen/logrus"
)

func main() {
	logrus.Info("Server is running...")

	http.HandleFunc("/", handlers.Hello) // each request calls handler

	logrus.Fatal(http.ListenAndServe("localhost:8000", nil))
}
