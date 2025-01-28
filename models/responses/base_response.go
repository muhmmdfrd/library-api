package responses

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(success bool, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(message string, data interface{}) BaseResponse {
	return newResponse(true, message, data)
}

func FailedResponse(message string) BaseResponse {
	return newResponse(false, message, nil)
}