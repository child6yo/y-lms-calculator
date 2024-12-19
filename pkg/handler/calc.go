package handler

import (
	"fmt"
	"net/http"
)

type RequestModel struct {
	Expression string `json:"name"`
}

type ResponseModel struct {
	Result string `json:"result"`
}

func (h *Handler) CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := h.services.Calculator.Calc("2+2*2")
	fmt.Fprint(w, n)
}
