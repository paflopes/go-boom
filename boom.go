// Package boom helps you to create and return HTTP errors easily
package boom

import (
	"encoding/json"
)

// Boom error structure. It implements the Error() method so
// you can return it as a regular error.
//
// If you don't want any data or message you can just omit them
// in the functions below
type Boom struct {
	StatusCode int         `json:"statusCode"`
	Err        string      `json:"error"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func (boom Boom) Error() string {
	errString, _ := json.Marshal(boom)
	return string(errString)
}

// IsServer is a utility method to check if the code is >= 500.
func (boom *Boom) IsServer() bool {
	return boom.StatusCode >= 500
}

// BadRequest returns a 400 Bad Request error
func BadRequest(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 400,
		Err:        "Bad Request",
		Data:       data,
		Message:    message,
	}
}

// Unauthorized returns a 401 Unauthorized error
func Unauthorized(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 401,
		Err:        "Unauthorized",
		Data:       data,
		Message:    message,
	}
}

// PaymentRequired returns a 402 Payment Required error
func PaymentRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 402,
		Err:        "Payment Required",
		Data:       data,
		Message:    message,
	}
}

// Forbidden returns a 403 Forbidden error
func Forbidden(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 403,
		Err:        "Forbidden",
		Data:       data,
		Message:    message,
	}
}

// NotFound returns a 404 Not Found error
func NotFound(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 404,
		Err:        "Not Found",
		Data:       data,
		Message:    message,
	}
}

// MethodNotAllowed returns a 405 Method Not Allowed error
func MethodNotAllowed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 405,
		Err:        "Method Not Allowed",
		Data:       data,
		Message:    message,
	}
}

// NotAcceptable returns a 406 Not Acceptable error
func NotAcceptable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 406,
		Err:        "Not Acceptable",
		Data:       data,
		Message:    message,
	}
}

// ProxyAuthRequired returns a 407 Proxy Auth Required error
func ProxyAuthRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 407,
		Err:        "Proxy Authentication Required",
		Data:       data,
		Message:    message,
	}
}

// ClientTimeout returns a 408 Client Timeout error
func ClientTimeout(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 408,
		Err:        "Request Time-out",
		Data:       data,
		Message:    message,
	}
}

// Conflict returns a 409 Conflict error
func Conflict(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 409,
		Err:        "Conflict",
		Data:       data,
		Message:    message,
	}
}

// ResourceGone returns a 410 Resource Gone error
func ResourceGone(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 410,
		Err:        "Gone",
		Data:       data,
		Message:    message,
	}
}

// LengthRequired returns a 411 Length Required error
func LengthRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 411,
		Err:        "Length Required",
		Data:       data,
		Message:    message,
	}
}

// PreconditionFailed returns a 412 Precondition Failed error
func PreconditionFailed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 412,
		Err:        "Precondition Failed",
		Data:       data,
		Message:    message,
	}
}

// EntityTooLarge returns a 413 Entity too Large error
func EntityTooLarge(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 413,
		Err:        "Request Entity Too Large",
		Data:       data,
		Message:    message,
	}
}

// URITooLong returns a 414 URI too Long error
func URITooLong(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 414,
		Err:        "Request-URI Too Large",
		Data:       data,
		Message:    message,
	}
}

// UnsupportedMediaType returns a 415 Unsupported Media Type error
func UnsupportedMediaType(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 415,
		Err:        "Unsupported Media Type",
		Data:       data,
		Message:    message,
	}
}

// RangeNotSatisfiable returns a 416 Requested Range Not Satisfiable error
func RangeNotSatisfiable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 416,
		Err:        "Requested Range Not Satisfiable",
		Data:       data,
		Message:    message,
	}
}

// ExpectationFailed returns a 417 Expectation Failed error
func ExpectationFailed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 417,
		Err:        "Expectation Failed",
		Data:       data,
		Message:    message,
	}
}

// Teapot returns a 418 Teapot error
func Teapot(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 418,
		Err:        "I'm a Teapot",
		Data:       data,
		Message:    message,
	}
}

// BadData returns a 422 Unprocessable Entity error
func BadData(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 422,
		Err:        "Unprocessable Entity",
		Data:       data,
		Message:    message,
	}
}

// Locked returns a 423 Locked error
func Locked(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 423,
		Err:        "Locked",
		Data:       data,
		Message:    message,
	}
}

// PreconditionRequired returns a 428 Precondition Required error
func PreconditionRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 428,
		Err:        "Precondition Required",
		Data:       data,
		Message:    message,
	}
}

// TooManyRequests returns a 429 Too Many Requests error
func TooManyRequests(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 429,
		Err:        "Too Many Requests",
		Data:       data,
		Message:    message,
	}
}

// Illegal returns a 451 Unavailable For Legal Reasons error
func Illegal(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 451,
		Err:        "Unavailable For Legal Reasons",
		Data:       data,
		Message:    message,
	}
}

// BadImplementation returns a 500 Internal Server Error error
func BadImplementation(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 500,
		Err:        "Internal Server Error",
		Data:       data,
		Message:    message,
	}
}

// NotImplemented returns a 501 Not Implemented error
func NotImplemented(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 501,
		Err:        "Not Implemented",
		Data:       data,
		Message:    message,
	}
}

// BadGateway returns a 502 Bad Gateway error
func BadGateway(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 502,
		Err:        "Bad Gateway",
		Data:       data,
		Message:    message,
	}
}

// ServerUnavailable returns a 503 Server Unavailable error
func ServerUnavailable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 503,
		Err:        "Service Unavailable",
		Data:       data,
		Message:    message,
	}
}

// GatewayTimeout returns a 504 Gateway Time-out error
func GatewayTimeout(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 504,
		Err:        "Gateway Time-out",
		Data:       data,
		Message:    message,
	}
}
