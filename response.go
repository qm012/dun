package ok

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// Response Return the complete information to the client
type Response struct {
	*StatusCode
	Data any `json:"data,omitempty" xml:"data,omitempty"` // Data object
}

// NewResponse The constructor of Response
func NewResponse(statusCode *StatusCode, data ...any) *Response {

	response := &Response{
		StatusCode: statusCode,
	}

	if len(data) > 0 && data[0] != nil && !reflect.ValueOf(data[0]).IsZero() {
		response.Data = data[0]
	}

	return response
}

// Success200 Return can carry data on success
//
//	dataChain[0] It must be an StatusCode object
func Success200(c *gin.Context, data ...any) {
	c.JSON(http.StatusOK, NewResponse(StatusCode200, data...))
}

// Success201 Return success0
func Success201(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// Failed200 http statusCode 200 result
//
//	dataChain[0] It must be an StatusCode object
func Failed200(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusOK, dataChain...)
}

// Failed400 Parameter error return, can carry a specific parameter error
//
//	dataChain[0] It must be an StatusCode object
func Failed400(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusBadRequest, dataChain...)
}

// Failed401 No access permission, identity authentication is required
//
//	dataChain[0] It must be an StatusCode object
func Failed401(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusUnauthorized, dataChain...)
}

// Failed402 Payment required
//
//	dataChain[0] It must be an StatusCode object
func Failed402(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusPaymentRequired, dataChain...)
}

// Failed403 Deny access
//
//	dataChain[0] It must be an StatusCode object
func Failed403(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusForbidden, dataChain...)
}

// Failed429 Too many requests
//
//	dataChain[0] It must be an StatusCode object
func Failed429(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusTooManyRequests, dataChain...)
}

// Failed500 internal server error
func Failed500(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusInternalServerError, dataChain...)
}

// Failed501 Not implemented
//
//	dataChain[0] It must be an StatusCode object
func Failed501(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusNotImplemented, dataChain...)
}

// Failed502 Bad gateway
//
//	dataChain[0] It must be an StatusCode object
func Failed502(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusBadGateway, dataChain...)
}

// Failed503 Service unavailable
//
//	dataChain[0] It must be an StatusCode object
func Failed503(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusServiceUnavailable, dataChain...)
}

// Failed504 Gateway timeout
// dataChain[0] It must be an StatusCode object
func Failed504(c *gin.Context, dataChain ...any) {
	Failed(c, http.StatusGatewayTimeout, dataChain...)
}

// Failed Custom return code
//
//	dataChain[0] It must be an StatusCode object
//	dataChain[1] You can return the specific error data corresponding to the error
func Failed(c *gin.Context, code int, dataChain ...any) {

	var (
		statusCode = StatusCode500() // default
		data       any               // result json data
	)

	length := len(dataChain)
	if length > 0 {
		var ok bool
		statusCode, ok = dataChain[0].(*StatusCode)
		if !ok {
			statusCode = StatusCode500()
		}
	}

	if length > 1 {
		data = dataChain[1]
	}
	c.JSON(code, NewResponse(statusCode, data))
}
