package response


type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	Error() string
	GetData() interface{}
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Err    string      `json:"error,omitempty"`
	Meta   *meta.Meta  `json:"meta,omitempty"`
}


