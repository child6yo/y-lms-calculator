package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"log/slog"

	"github.com/child6yo/y-lms-calculator/pkg/service"
)

type RequestModel struct {
	Expression string `json:"expression"`
}

type ResponseModel struct {
	Result string `json:"result"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestModel

	data, err := io.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		httpNewError(w, 500, "Internal server error")
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &req)
	if err != nil {
		httpNewError(w, 422, "Expression is not valid")
		return
	}

	res, err := service.Calc(req.Expression)
	if err != nil {
		httpNewError(w, 422, "Expression is not valid")
		return
	}
	response := ResponseModel{Result: fmt.Sprintf("%v", res)}
	responseData, err := json.MarshalIndent(response, "", " ")
	if err != nil { 
		httpNewError(w, 500, "Internal server error")
		return 
	}

	w.Header().Set("Content-Type", "application/json") 
	w.Write(responseData)
}

type ErrorModel struct {
	Error string `json:"error"`
}

func httpNewError(w http.ResponseWriter, statusCode int, message string) {
	slog.Error(message)

	response := ErrorModel{Error: message}
	responseData, _ := json.MarshalIndent(response, "", " ")

	http.Error(w, string(responseData), statusCode)
}