package appconstants

import "errors"

var (
	DuplicateEmailError = errors.New("error duplicate email")
	WrongKeyLoginError  = errors.New("can not login wrong email or password")
)

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
