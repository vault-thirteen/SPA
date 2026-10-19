package server

type ProxyError struct {
	StatusCode int
	Message    string
}

func NewProxyError(statusCode int, message string) *ProxyError {
	return &ProxyError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (pe *ProxyError) Error() string {
	return pe.Message
}
