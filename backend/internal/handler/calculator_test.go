package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "adds two numbers",
			method: http.MethodPost,
			body: `{
				"operation": "add",
				"operand1": 2,
				"operand2": 3
			}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":5}`,
		},
		{
			name:   "subtracts two numbers",
			method: http.MethodPost,
			body: `{
				"operation": "subtract",
				"operand1": 10,
				"operand2": 3
			}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":7}`,
		},
		{
			name:   "multiplies two numbers",
			method: http.MethodPost,
			body: `{
				"operation": "multiply",
				"operand1": 5,
				"operand2": 4
			}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":20}`,
		},
		{
			name:   "divides two numbers",
			method: http.MethodPost,
			body: `{
				"operation": "divide",
				"operand1": 10,
				"operand2": 2
			}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":5}`,
		},
		{
			name:   "returns bad request for division by zero",
			method: http.MethodPost,
			body: `{
				"operation": "divide",
				"operand1": 10,
				"operand2": 0
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"division by zero is not allowed"}`,
		},
		{
			name:   "returns bad request when second operand is missing",
			method: http.MethodPost,
			body: `{
				"operation": "add",
				"operand1": 10
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"missing second operand"}`,
		},
		{
			name:   "returns bad request for unsupported operation",
			method: http.MethodPost,
			body: `{
				"operation": "modulo",
				"operand1": 10,
				"operand2": 3
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"unsupported operation"}`,
		},
		{
			name:   "returns bad request for invalid json",
			method: http.MethodPost,
			body: `{
				"operation": "add",
				"operand1":
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid request body"}`,
		},
		{
			name:           "returns method not allowed for get",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   `{"error":"method not allowed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				tt.method,
				"/api/calculate",
				strings.NewReader(tt.body),
			)

			recorder := httptest.NewRecorder()

			Calculate(recorder, req)

			if recorder.Code != tt.expectedStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.expectedStatus,
					recorder.Code,
				)
			}

			actualBody := strings.TrimSpace(recorder.Body.String())

			if actualBody != tt.expectedBody {
				t.Errorf(
					"expected body %s, got %s",
					tt.expectedBody,
					actualBody,
				)
			}
		})
	}
}
