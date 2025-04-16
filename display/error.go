package display

import (
	"errors"
	"net/http"
)

type CustomError struct {
	CodeErr    int    `json:"codeErr"`
	MessageErr string `json:"messageErr"`
}

type ErrorResponse struct {
	MessageErrRes string `json:"messageErrRes"`
}

func NewCustomErrorDisp(codeError int, messageError string) *CustomError {
	return &CustomError{
		CodeErr:    codeError,
		MessageErr: messageError,
	}
}

func (c *CustomError) ErrorDisp() string {

	return c.MessageErr
}

func (c *CustomError) ConvertToErrorResponse() ErrorResponse {
	return ErrorResponse{
		MessageErrRes: c.MessageErr,
	}
}

var (
	ErrorWrongCredentialsLogin = errors.New("wrong email or password")
	ErrorWrongCredentials      = NewCustomErrorDisp(http.StatusUnauthorized, "wrong credentials")

	ErrorInvalidBody        = NewCustomErrorDisp(http.StatusBadRequest, "invalid body")
	ErrorInvalidParamID     = NewCustomErrorDisp(http.StatusBadRequest, "invalid param id")
	ErrorUnathorized        = NewCustomErrorDisp(http.StatusUnauthorized, "you are not authorized")
	ErrorBearerTokenInvalid = NewCustomErrorDisp(http.StatusUnauthorized, "bearer token is invalid")
)
