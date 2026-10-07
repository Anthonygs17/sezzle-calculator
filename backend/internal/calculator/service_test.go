package calculator

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name          string
		operation     string
		operand1      float64
		operand2      float64
		expected      float64
		expectedError error
	}{
		{
			name:      "adds two numbers",
			operation: "add",
			operand1:  2,
			operand2:  3,
			expected:  5,
		},
		{
			name:      "subtracts two numbers",
			operation: "subtract",
			operand1:  10,
			operand2:  3,
			expected:  7,
		},
		{
			name:      "multiplies two numbers",
			operation: "multiply",
			operand1:  5,
			operand2:  4,
			expected:  20,
		},
		{
			name:      "divides two numbers",
			operation: "divide",
			operand1:  10,
			operand2:  2,
			expected:  5,
		},
		{
			name:          "returns error when dividing by zero",
			operation:     "divide",
			operand1:      10,
			operand2:      0,
			expectedError: ErrDivisionByZero,
		},
		{
			name:          "returns error for unsupported operation",
			operation:     "modulo",
			operand1:      10,
			operand2:      3,
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
