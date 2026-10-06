package handler

type CalculateRequest struct {
	Operation string   `json:"operation"`
	Operand1  float64  `json:"operand1"`
	Operand2  *float64 `json:"operand2"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
