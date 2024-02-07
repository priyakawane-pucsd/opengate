package dto

type PingResponse struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"message"`
}
