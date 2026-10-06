package calculator

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	operand2Three := 3.0
	operand2Four := 4.0
	operand2Two := 2.0
	operand2Zero := 0.0

	tests := []struct {
		name          string
		operation     string
		operand1      float64
		operand2      *float64
		expected      float64
		expectedError error
	}{
		{
			name:      "adds two numbers",
			operation: "add",
			operand1:  2,
			operand2:  &operand2Three,
			expected:  5,
		},
		{
			name:      "subtracts two numbers",
			operation: "subtract",
			operand1:  10,
			operand2:  &operand2Three,
			expected:  7,
		},
		{
			name:      "multiplies two numbers",
			operation: "multiply",
			operand1:  5,
			operand2:  &operand2Four,
			expected:  20,
		},
		{
			name:      "divides two numbers",
			operation: "divide",
			operand1:  10,
			operand2:  &operand2Two,
			expected:  5,
		},
		{
			name:          "returns error when dividing by zero",
			operation:     "divide",
			operand1:      10,
			operand2:      &operand2Zero,
			expectedError: ErrDivisionByZero,
		},
		{
			name:          "returns error when second operand is missing",
			operation:     "add",
			operand1:      10,
			operand2:      nil,
			expectedError: ErrMissingOperand,
		},
		{
			name:          "returns error for unsupported operation",
			operation:     "modulo",
			operand1:      10,
			operand2:      &operand2Two,
			expectedError: ErrUnsupportedOperation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(
				tt.operation,
				tt.operand1,
				tt.operand2,
			)

			if tt.expectedError != nil {
				if !errors.Is(err, tt.expectedError) {
					t.Fatalf(
						"expected error %v, got %v",
						tt.expectedError,
						err,
					)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf(
					"expected result %v, got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}
