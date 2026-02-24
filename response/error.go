package response

import (
	"net/http"
)


type ErrorResponse struct {
	Status int         `json:"status"`
	ErrorMessage      string      `json:"error,omitempty"`

}


func buildResponseError(message string, status int) Response {
	return &ErrorResponse{
		Status:  status,
		ErrorMessage:    message,
	}
}	


func BadRequest(message string) Response {
	return buildResponseError(message, http.StatusBadRequest)
}

func InternalServerError(message string) Response {
	return buildResponseError(message, http.StatusInternalServerError)
}


func (e *ErrorResponse) StatusCode() int {
	return e.Status
}


func (e *ErrorResponse) GetBody() ([]byte, error) {
	return []byte(e.ErrorMessage), nil
}

func (e *ErrorResponse) Error() string {
	return e.ErrorMessage
}

func (e *ErrorResponse) GetData() interface{} {
	return interface{}(nil)
}