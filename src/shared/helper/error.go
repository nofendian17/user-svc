package helper

type ApplicationError struct {
	Code    string
	Message string
}

func NewError(code string, message string) error {
	return &ApplicationError{
		Code:    code,
		Message: message,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) SetErrorCode(errorCode string) *ApplicationError {
	e.Code = errorCode
	return e
}

func (e *ApplicationError) SetMessage(message string) *ApplicationError {
	e.Message = message
	return e
}
