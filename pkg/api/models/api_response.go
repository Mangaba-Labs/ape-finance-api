package models

// ApiResponse definition
type ApiResponse struct {
	HttpCode int
	Message  string
	Status   string
}

// Set all fields of ApiResponse
func (a *ApiResponse) Set(status string, message string, httpCode int) {
	a.HttpCode = httpCode
	a.Message = message
	a.Status = status
}
