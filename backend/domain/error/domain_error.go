package domain

import (
	"net/http"
)

type Error interface {
	error
	Status() int
	Message() string
}

type ErrorFn func(int, string) Error

type BaseError struct {
	status  int
	message string
}

func NewError(status int, message string) Error {
	return BaseError{
		status:  status,
		message: message,
	}
}

func (e BaseError) Error() string {
	return e.message
}

func (e BaseError) Status() int {
	return e.status
}

func (e BaseError) Message() string {
	return e.message
}

func BadRequest(message string) Error {
	return NewError(http.StatusBadRequest, message)
}

func InternalServerError(message string) Error {
	return NewError(http.StatusInternalServerError, message)
}
