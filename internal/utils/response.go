// internal/utils/response.go
package utils

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}, message string) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(err string) Response {
	return Response{
		Status: "error",
		Error:  err,
	}
}