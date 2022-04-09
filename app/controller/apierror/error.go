package apierror

import "net/http"

type ApiError interface {
	Message() string
	Status() int
}

type ApiErr struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func BadRequestError(message string) ApiErr {
	return ApiErr{message, http.StatusBadRequest}
}

func (e ApiErr) Message() string {
	return e.ErrorMessage
}
