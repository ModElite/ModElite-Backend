package domain

type Response struct {
	SUCCESS bool        `json:"success"`
	MESSAGE string      `json:"message"`
	DATA    interface{} `json:"data,omitempty"`
	ERROR   error       `json:"error,omitempty"`
}
