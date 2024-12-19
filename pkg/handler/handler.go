package handler

import "net/http"

func InitEndpoints() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalculatorHandler)

	return mux
}