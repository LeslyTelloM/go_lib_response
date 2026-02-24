package response

import (
	"net/http"
)"github.com/LeslyTelloM/go_meta/meta"


type SuccessResponse struct {
	Status int         `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message,omitempty"`
	Meta 	 *meta.Meta       `json:"meta,omitempty"`

}


func buildResponse(message string,data interface{}, status int, meta *meta.Meta) Response {
	return &Response{
		Status:  status,
		Data:    data,
		Message: message,
		Meta: meta
	}
}	


func Created(message string, data interface{}, meta *meta.Meta) Response {
	return buildResponse(message, data, http.StatusCreated, meta)
}

func OK(message string, data interface{}, meta *meta.Meta) Response {
	return buildResponse(message, data, http.StatusOK, meta)
}



func (s *SuccessResponse) StatusCode() int {
	return s.Status
}


func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s.Data)
}

func (s *SuccessResponse) Error() string {
	return s.Message
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
