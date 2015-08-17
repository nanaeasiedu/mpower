package mpower

type Response struct {
	ResponseText string `json:"response_text,omitempty"`
	ResponseCode string `json:"response_code,omitempty"`
	Description  string `json:"description,omitempty"`
}
