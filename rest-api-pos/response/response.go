package response

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(code int, message string) APIResponse {
	return APIResponse{
		Code:    code,
		Message: message,
	}
}
