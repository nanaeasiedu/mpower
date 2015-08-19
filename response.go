package mpower

// Response - almost all mpower JSON responses contain all these fields
// This struct is embeded in other structs
type Response struct {
	ResponseText string `json:"response_text,omitempty"`
	ResponseCode string `json:"response_code,omitempty"`
	Description  string `json:"description,omitempty"`
}
