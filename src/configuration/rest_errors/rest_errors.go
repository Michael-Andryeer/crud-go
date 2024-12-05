package rest_errors

import "net/http"

type RestError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`   // Código da requisição para retornar ao cliente
	Causes  []Causes `json:"causes"` //Causas de erro

}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewRestError(message, err string, code int, causes []Causes) *RestError {
	return &RestError{Message: message,
		Err:    err,
		Code:   code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{Message: message,
		Err:  "bad request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{Message: message,
		Err:    "bad request",
		Code:   http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbbidenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
