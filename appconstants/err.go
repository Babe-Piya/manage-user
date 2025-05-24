package appconstants

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    1,
		Message: err.Error(),
	}
}
