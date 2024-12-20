package handler

import "net/http"

func InitEndpoints() *http.ServeMux {
	mux := http.NewServeMux()
	logger := Logger(CalculatorHandler)
	mux.Handle("/api/v1/calculate", logger)

	return mux
}