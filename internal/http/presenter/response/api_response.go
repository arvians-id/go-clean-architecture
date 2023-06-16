package response

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type ErrorValidationResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Error  []*ErrorResponse `json:"error"`
}
