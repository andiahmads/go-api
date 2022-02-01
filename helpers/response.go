package helpers

import "strings"

//response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

type CustomeResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used  when data dosn't want to be null on json'
type EmptyObj struct{}

//BuildResponse method is to inject data value do dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildResponse method is to inject data value do dynamic errors response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	//pecah array error
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res

}
