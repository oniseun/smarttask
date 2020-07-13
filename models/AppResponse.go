package models


type AppResponse struct {
	ResponseCode    int    `json:"responseCode"`
	Message  string    `json:"message"`
	Data interface{} `json:"data"`
}
