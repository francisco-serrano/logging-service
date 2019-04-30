package utils

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code        string
	Description string
}

const (
	ErrInvalidData string = "err_invalid_data"
)

func ErrorInvalidData(err error) *Error {
	return &Error{Code: ErrInvalidData, Description: fmt.Sprintf("%s", err)}
}

func GetStatusErrorCode(error *Error) int {
	switch error.Code {
	case ErrInvalidData:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
