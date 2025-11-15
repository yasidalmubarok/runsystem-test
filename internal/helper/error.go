package helper

import "net/http"

type ErrorInterface interface {
	Status() int
	Message() string
	Error() string
}

type ErrorStruct struct {
	ErrStatus  int
	ErrMessage string
	ErrError   string
}

func (e *ErrorStruct) Status() int {
	return e.ErrStatus
}

func (e *ErrorStruct) Message() string {
	return e.ErrMessage
}

func (e *ErrorStruct) Error() string {
	return e.ErrError
}

func NewInternalServerError(message string) ErrorInterface {
	return &ErrorStruct{
		ErrStatus:  500,
		ErrMessage: message,
		ErrError:   "Internal Server Error",
	}
}

func NewBadRequestError(message string) ErrorInterface {
	return &ErrorStruct{
		ErrStatus:  400,
		ErrMessage: message,
		ErrError:   "Bad Request",
	}
}

func NewNotFoundError(message string) ErrorInterface {
	return &ErrorStruct{
		ErrStatus:  404,
		ErrMessage: message,
		ErrError:   "Not Found",
	}
}

func NewConflictError(message string) ErrorInterface {
	return &ErrorStruct{
		ErrStatus:  409,
		ErrMessage: message,
		ErrError:   "Conflict",
	}
}

func NewUnprocessableEntityError(message string) ErrorInterface {
	return &ErrorStruct{
		ErrStatus: http.StatusUnprocessableEntity,
		ErrMessage: message,
		ErrError:   "Invalid Request Body",
	}
}