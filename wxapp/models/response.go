package models

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	Success         = 0
	Fail            = 100
	ArgumentInvalid = 400
)

func SuccessResponse(data interface{}) BaseResponse {
	resp := BaseResponse{}
	resp.Code = Success
	resp.Data = data

	return resp
}

func FailResponseWithCode(code int, message string) BaseResponse {
	resp := BaseResponse{}
	resp.Code = code
	resp.Message = message

	return resp
}

func FailResponse(message string) BaseResponse {
	resp := BaseResponse{}
	resp.Code = Fail
	resp.Message = message

	return resp
}
