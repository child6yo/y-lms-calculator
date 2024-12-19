package handler

import (
	"net/http"

	"github.com/child6yo/y-lms-calculator/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitEndpoints() (*http.ServeMux) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", h.CalculatorHandler)

	return mux
}