package web

type StatusResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
	Success    bool        `json:"success,omitempty"`
}

func CreateResponse(statusCode int, errorMessage interface{}, data interface{}, success bool) StatusResponse {
	return StatusResponse{
		StatusCode: statusCode,
		Errors:     errorMessage,
		Data:       data,
		Success:    success,
	}
}
