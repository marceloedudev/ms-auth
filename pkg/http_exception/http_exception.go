package http_exception

import (
	"ms-auth/pkg/utils"
	"net/http"
	"time"
)

type HttpException struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHttpException(status int, message string) *HttpException {
	return &HttpException{
		Message: message,
		Status:  status,
	}
}

func InternalServerException(err error) *HttpException {
	return NewHttpException(http.StatusInternalServerError, utils.ValidatorError(err))
}

func BadRequestException(err error) *HttpException {
	return NewHttpException(http.StatusBadRequest, utils.ValidatorError(err))
}

func NotFoundExeception(err error) *HttpException {
	return NewHttpException(http.StatusNotFound, utils.ValidatorError(err))
}

func UnauthorizedExeception(err error) *HttpException {
	return NewHttpException(http.StatusUnauthorized, utils.ValidatorError(err))
}
