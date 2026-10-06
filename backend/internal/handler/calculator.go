package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Anthonygs17/sezzle-calculator/backend/internal/calculator"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "method not allowed",
		})

		return
	}

	var req CalculateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "invalid request body",
		})

		return
	}

	result, err := calculator.Calculate(req.Operation, req.Operand1, req.Operand2)

	if err != nil {
		switch {
		case errors.Is(err, calculator.ErrMissingOperand),
			errors.Is(err, calculator.ErrDivisionByZero),
			errors.Is(err, calculator.ErrUnsupportedOperation):
			w.WriteHeader(http.StatusBadRequest)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(CalculateResponse{
		Result: result,
	})
}
