package calculator

import "errors"

var (
	ErrMissingOperand       = errors.New("missing second operand")
	ErrDivisionByZero       = errors.New("division by zero is not allowed")
	ErrUnsupportedOperation = errors.New("unsupported operation")
)

func Calculate(
	operation string,
	operand1 float64,
	operand2 float64,
) (float64, error) {
	switch operation {
	case "add":
		return operand1 + operand2, nil

	case "subtract":
		return operand1 - operand2, nil

	case "multiply":
		return operand1 * operand2, nil

	case "divide":
		if operand2 == 0 {
			return 0, ErrDivisionByZero
		}
		return operand1 / operand2, nil

	default:
		return 0, ErrUnsupportedOperation
	}
}
