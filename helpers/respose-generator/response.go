package resposegenerator

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(success bool, message string, data interface{}) *Response {
	return &Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}
