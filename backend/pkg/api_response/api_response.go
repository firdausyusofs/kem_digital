package api_response

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}

func MakeSuccessResponse(statusCode int, message string, data interface{}) (int, APIResponse) {
	return statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func MakeErrorResponse(statusCode int, message string, err *string) (int, APIResponse) {
	return statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
