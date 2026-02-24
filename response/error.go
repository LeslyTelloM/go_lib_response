package response

import (
	"net/http"
)"github.com/LeslyTelloM/go_meta/meta"


type ErrorResponse struct {
	Status int         `json:"status"`
	Error      string      `json:"error,omitempty"`

}


func buildResponse(message string, status int) Response {
	return &Response{
		Status:  status,
		Error:    message,
	}
}	


func BadRequest(message string) Response {
	return buildResponse(message, http.StatusBadRequest)
}

func InternalServerError(message string) Response {
	return buildResponse(message, http.StatusInternalServerError)
}


func (e *ErrorResponse) StatusCode() int {
	return e.Status
}


func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e.Error)
}

func (e *ErrorResponse) Error() string {
	return e.Error
}

func (e *ErrorResponse) GetData() interface{} {
	return e.Error
}