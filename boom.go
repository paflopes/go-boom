package boom

import (
	"encoding/json"
)

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

func (boom *Boom) IsServer() bool {
	return boom.StatusCode <= 500
}

func BadRequest(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 400,
		Err:        "Bad Request",
		Data:       data,
		Message:    message,
	}
}

func Unauthorized(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 401,
		Err:        "Unauthorized",
		Data:       data,
		Message:    message,
	}
}

func PaymentRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 402,
		Err:        "Payment Required",
		Data:       data,
		Message:    message,
	}
}

func Forbidden(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 403,
		Err:        "Forbidden",
		Data:       data,
		Message:    message,
	}
}

func NotFound(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 404,
		Err:        "Not Found",
		Data:       data,
		Message:    message,
	}
}

func MethodNotAllowed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 405,
		Err:        "Method Not Allowed",
		Data:       data,
		Message:    message,
	}
}

func NotAcceptable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 406,
		Err:        "Not Acceptable",
		Data:       data,
		Message:    message,
	}
}

func ProxyAuthRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 407,
		Err:        "Proxy Authentication Required",
		Data:       data,
		Message:    message,
	}
}

func ClientTimeout(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 408,
		Err:        "Request Time-out",
		Data:       data,
		Message:    message,
	}
}

func Conflict(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 409,
		Err:        "Conflict",
		Data:       data,
		Message:    message,
	}
}

func ResourceGone(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 410,
		Err:        "Gone",
		Data:       data,
		Message:    message,
	}
}

func LengthRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 411,
		Err:        "Length Required",
		Data:       data,
		Message:    message,
	}
}

func PreconditionFailed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 412,
		Err:        "Precondition Failed",
		Data:       data,
		Message:    message,
	}
}

func EntityTooLarge(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 413,
		Err:        "Request Entity Too Large",
		Data:       data,
		Message:    message,
	}
}

func UriTooLong(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 414,
		Err:        "Request-URI Too Large",
		Data:       data,
		Message:    message,
	}
}

func UnsupportedMediaType(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 415,
		Err:        "Unsupported Media Type",
		Data:       data,
		Message:    message,
	}
}

func RangeNotSatisfiable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 416,
		Err:        "Requested Range Not Satisfiable",
		Data:       data,
		Message:    message,
	}
}

func ExpectationFailed(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 417,
		Err:        "Expectation Failed",
		Data:       data,
		Message:    message,
	}
}

func Teapot(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 418,
		Err:        "I'm a Teapot",
		Data:       data,
		Message:    message,
	}
}

func BadData(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 422,
		Err:        "Unprocessable Entity",
		Data:       data,
		Message:    message,
	}
}

func Locked(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 423,
		Err:        "Locked",
		Data:       data,
		Message:    message,
	}
}

func PreconditionRequired(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 428,
		Err:        "Precondition Required",
		Data:       data,
		Message:    message,
	}
}

func TooManyRequests(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 429,
		Err:        "Too Many Requests",
		Data:       data,
		Message:    message,
	}
}

func Illegal(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 451,
		Err:        "Unavailable For Legal Reasons",
		Data:       data,
		Message:    message,
	}
}

func BadImplementation(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 500,
		Err:        "Internal Server Error",
		Data:       data,
		Message:    message,
	}
}

func NotImplemented(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 501,
		Err:        "Not Implemented",
		Data:       data,
		Message:    message,
	}
}

func BadGateway(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 502,
		Err:        "Bad Gateway",
		Data:       data,
		Message:    message,
	}
}

func ServerUnavailable(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 503,
		Err:        "Service Unavailable",
		Data:       data,
		Message:    message,
	}
}

func GatewayTimeout(message string, data interface{}) *Boom {
	return &Boom{
		StatusCode: 504,
		Err:        "Gateway Time-out",
		Data:       data,
		Message:    message,
	}
}
