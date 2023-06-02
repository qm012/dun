package dun

import (
	"fmt"
)

var (
	StatusCodeSuccess         = NewStatusCode(200, "success")                                         // success case status code
	StatusCodeSystemError     = NewStatusCode(500, "system error")                                    // default status code
	StatusCodeExternalService = NewStatusCode(10001, "external services are unavailable")             // Third party service：A situation in which external services are unavailable
	StatusCodeInternalService = NewStatusCode(10002, "internal services are unavailable")             // Other services in the microservice：An internal service is unavailable
	StatusCodeSystemBusy      = NewStatusCode(10003, "the system is busy, please try again later...") // When the system is busy
	StatusCodeApiNotOpened    = NewStatusCode(10004, "this API is not open yet, so stay tuned!")
	StatusCodeApiDeprecated   = NewStatusCode(10005, "this API is deprecated and is no longer used!")
	StatusCodeDataNotFound    = NewStatusCode(10006, "data error")
)

// StatusCode Custom return status code messages
type StatusCode struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

// NewStatusCode The constructor of StatusCode
func NewStatusCode(code int, message string) *StatusCode {
	return &StatusCode{Code: code, Message: message}
}

// Error impl error
func (s *StatusCode) Error() string {
	return fmt.Sprintf("code:%d message:%s", s.Code, s.Message)
}
