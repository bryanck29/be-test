package utils

import (
	"net/http"

	"github.com/bryanck29/be-test/internal/constant"
)

// ErrorMapping maps error and return its corresponding http status and the err obj
func ErrorMapping(err error) (int, error) {
	switch err {
	case constant.ErrParsingRequest:
		return constant.CommonErrorMap[constant.ErrParsingRequest], constant.ErrParsingRequest
	case constant.ErrNoAuth:
		return constant.CommonErrorMap[constant.ErrNoAuth], constant.ErrNoAuth
	case constant.ErrInvalidAuth:
		return constant.CommonErrorMap[constant.ErrInvalidAuth], constant.ErrInvalidAuth
	case constant.ErrInvalidSession:
		return constant.CommonErrorMap[constant.ErrInvalidSession], constant.ErrInvalidSession
	case constant.ErrDataNotFound:
		return constant.CommonErrorMap[constant.ErrDataNotFound], constant.ErrDataNotFound
	case constant.ErrDataExists:
		return constant.CommonErrorMap[constant.ErrDataExists], constant.ErrDataExists
	case constant.ErrInvalidLoginCredential:
		return constant.CommonErrorMap[constant.ErrInvalidLoginCredential], constant.ErrInvalidLoginCredential
	case constant.ErrInvalidAuthToken:
		return constant.CommonErrorMap[constant.ErrInvalidAuthToken], constant.ErrInvalidAuthToken
	case constant.ErrInvalidAccess:
		return constant.CommonErrorMap[constant.ErrInvalidAccess], constant.ErrInvalidAccess
	default:
		return http.StatusInternalServerError, err
	}
}
