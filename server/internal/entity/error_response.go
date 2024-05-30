package entity

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func BuildErrorResponse(statusCode int, err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}
