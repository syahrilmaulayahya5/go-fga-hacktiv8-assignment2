package message

import (
	"net/http"
	"time"
)

type BaseResponse struct {
	Meta struct {
		Status    int        `json:"status"`
		Message   string     `json:"message"`
		StartTime *time.Time `json:"start_time,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) BaseResponse {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data = data
	return response
}

func NewErrorResponse(status int, message string) BaseResponse {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = message
	response.Data = nil
	return response
}
