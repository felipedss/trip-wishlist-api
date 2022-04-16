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

func BadRequestErrorQueryParms() ApiErr {
	return BadRequestError("Error to build query-params")
}

func BadRequestErrorCallRemote() ApiErr {
	return BadRequestError("Error to call a remote re")
}

func BadRequestErrorBindJson() ApiErr {
	return BadRequestError("Error to bind json")
}

func (e ApiErr) Message() string {
	return e.ErrorMessage
}
