package main

import (
	"net/http"

	"github.com/child6yo/y-lms-calculator/pkg/handler"
	"github.com/child6yo/y-lms-calculator/pkg/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitEndpoints()

	http.ListenAndServe(":8000", router)
}