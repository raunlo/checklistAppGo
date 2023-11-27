package domain

type Error interface {
	Error() string
	ResponseCode() int
}

type error struct {
	errorMessage string
	responseCode int
}

func (e *error) Error() string {
	return e.errorMessage
}

func (e *error) ResponseCode() int {
	return e.responseCode
}

func NewError(errorMessage string, responseCode int) Error {
	return &error{
		errorMessage: errorMessage,
		responseCode: responseCode,
	}
}
