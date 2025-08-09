package model

type ErrorResponse struct {
	Error APIError `json:"error"`
}

type APIError struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Details *ErrorDetails `json:"details"`
	TraceID string        `json:"trace_id"`
}

type ErrorDetails struct {
	Resource string `json:"resource"`
	Id       *int   `json:"id"`
}
