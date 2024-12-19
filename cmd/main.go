package main

import (
	"net/http"

	"github.com/child6yo/y-lms-calculator/pkg/handler"
)

func main() {

	router := handler.InitEndpoints()

	http.ListenAndServe(":8000", router)
}