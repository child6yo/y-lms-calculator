package handler

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"net/http/httptest"
)

type TestResponse struct { 
	Result string `json:"result,omitempty"` 
	Error string `json:"error,omitempty"` 
}

func TestCalcultatorHandler2(t *testing.T) {
	testCases := []struct {
		requestBody *strings.Reader
		expectBody TestResponse
		expectStatus int
	}{
		{
			requestBody: strings.NewReader(`{"expression": "2+2*2"}`),
			expectBody: TestResponse{Result: "6"},
			expectStatus: 200,
		},
		{
			requestBody: strings.NewReader(`{"expression": "(2+2)*2"}`),
			expectBody: TestResponse{Result: "8"},
			expectStatus: 200,
		},
		{
			requestBody: strings.NewReader(`{"expression": "2/0"}`),
			expectBody: TestResponse{Error: "Expression is not valid"},
			expectStatus: 422,
		},
		{
			requestBody: strings.NewReader(`{"expression": 2+2}`),
			expectBody: TestResponse{Error: "Expression is not valid"},
			expectStatus: 422,
		},
		{
			requestBody: strings.NewReader(""),
			expectBody: TestResponse{Error: "Internal server error"},
			expectStatus: 500,
		},
	}


	for i, tc := range testCases {
		req := httptest.NewRequest("POST", "http://localhost:8000/api/v1/calculate", tc.requestBody)
		w := httptest.NewRecorder()
		CalculatorHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("test %d: error reading response body: %v", i+1, err)
		}

		var c TestResponse
		json.Unmarshal(body, &c)


		if c != tc.expectBody {
			t.Errorf("test %d failed: result: %s, expected: %s", i+1, c, tc.expectBody)
		}
		if resp.StatusCode != tc.expectStatus {
			t.Errorf("test %d failed: result: %d, expected: %d", i+1, resp.StatusCode, tc.expectStatus)
		}
	}
}