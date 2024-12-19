package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestModel struct {
	Expression string `json:"expression"`
}

type ResponseModel struct {
	Result string `json:"result"`
}

func (h *Handler) CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var s RequestModel

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &s)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	res, err := h.services.Calculator.Calc(s.Expression)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	response := ResponseModel{Result: fmt.Sprintf("%v", res)}
	responseData, err := json.Marshal(response)
	if err != nil { 
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json") 
	w.Write(responseData)
}