package model

// HTTPResponse is a custom HTTP response template
type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
