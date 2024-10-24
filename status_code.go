package okk

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	StatusCode200             = NewStatusCode(http.StatusOK, errors.New("success"))                            // success case status code
	StatusCode404             = NewStatusCode(http.StatusNotFound, errors.New("route not found"))              // route not found case status code
	StatusCodeExternalService = NewStatusCode(10001, errors.New("external services are unavailable"))          // Third party service：A situation in which external services are unavailable
	StatusCodeInternalService = NewStatusCode(10002, errors.New("internal services are unavailable"))          // Other services in the microservice：An internal service is unavailable
	StatusCodeSystemBusy      = NewStatusCode(10003, errors.New("the system is busy, please try again later")) // When the system is busy
	StatusCodeApiNotOpened    = NewStatusCode(10004, errors.New("this API is not open yet, so stay tuned"))
	StatusCodeApiDeprecated   = NewStatusCode(10005, errors.New("this API is deprecated and is no longer used"))
	StatusCodeDataNotFound    = NewStatusCode(10006, errors.New("data error"))
)

// StatusCode Custom return status code msg
type StatusCode struct {
	Code int    `json:"code" xml:"code"`
	Msg  string `json:"msg" xml:"msg"`
}

// NewStatusCode The constructor of StatusCode
func NewStatusCode(code int, err error) *StatusCode {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &StatusCode{Code: code, Msg: errMsg}
}

func StatusCode400(errs ...error) *StatusCode {
	var tempErr = errors.New("parameter error")
	if len(errs) > 0 {
		tempErr = errs[0]
	}
	return NewStatusCode(http.StatusBadRequest, tempErr)
}

func StatusCode500(errs ...error) *StatusCode {
	var tempErr = errors.New("system error")
	if len(errs) > 0 {
		tempErr = errs[0]
	}
	return NewStatusCode(500, tempErr)
}

// Error impl error
func (s *StatusCode) Error() string {
	return fmt.Sprintf("code:%d msg:%s", s.Code, s.Msg)
}
